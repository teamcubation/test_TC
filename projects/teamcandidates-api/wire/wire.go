//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	gorm "github.com/teamcubation/teamcandidates/pkg/databases/sql/gorm"
	mdw "github.com/teamcubation/teamcandidates/pkg/http/middlewares/gin"
	ginsrv "github.com/teamcubation/teamcandidates/pkg/http/servers/gin"

	jwt "github.com/teamcubation/teamcandidates/pkg/authe/jwt/v5"
	rabbit "github.com/teamcubation/teamcandidates/pkg/brokers/rabbitmq/amqp091/producer"
	redis "github.com/teamcubation/teamcandidates/pkg/databases/cache/redis/v8"
	cass "github.com/teamcubation/teamcandidates/pkg/databases/nosql/cassandra/gocql"
	mongo "github.com/teamcubation/teamcandidates/pkg/databases/nosql/mongodb/mongo-driver"
	pg "github.com/teamcubation/teamcandidates/pkg/databases/sql/postgresql/pgxpool"
	resty "github.com/teamcubation/teamcandidates/pkg/http/clients/resty"
	smtp "github.com/teamcubation/teamcandidates/pkg/notification/smtp"
	ws "github.com/teamcubation/teamcandidates/pkg/websocket/gorilla"

	assessment "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/assessment"
	authe "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/authe"
	browserevent "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/browser-events"
	candidate "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/candidate"
	category "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/category"
	config "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/config"
	event "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/event"
	group "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/group"
	item "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/item"
	macrocategory "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/macrocategory"
	notification "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/notification"
	person "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/person"
	supplier "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/supplier"
	tweet "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/tweet"
	user "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/user"
)

// Dependencies reúne todas las dependencias de la aplicación que se inyectan con Wire.
type Dependencies struct {
	ConfigLoader        config.Loader
	GinServer           ginsrv.Server
	GormRepository      gorm.Repository
	MongoRepository     mongo.Repository
	PostgresRepository  pg.Repository
	RedisCache          redis.Cache
	JwtService          jwt.Service
	RestyClient         resty.Client
	SmtpService         smtp.Service
	RabbitProducer      rabbit.Producer
	CassandraRepository cass.Repository
	WebSocket           ws.Upgrader

	Middlewares *mdw.Middlewares

	PersonHandler          *person.Handler
	GroupHandler           *group.Handler
	EventHandler           *event.Handler
	UserHandler            *user.Handler
	AssessmentHandler      *assessment.Handler
	CandidateHandler       *candidate.Handler
	BrowserEventsHandler   *browserevent.Handler
	BrowserEventsWebSocket browserevent.WebSocket
	AutheHandler           *authe.Handler
	NotificationHandler    *notification.Handler
	TweetHandler           *tweet.Handler
	ItemHandler            *item.Handler
	CategoryHandler        *category.Handler
	MacroCategoryHandler   *macrocategory.Handler
	SupplierHandler        *supplier.Handler

	// Para pruebas
	PersonUseCases person.UseCases
	UserUseCases   user.UseCases
	TweetUseCases  tweet.UseCases
	ItemUseCases   item.UseCases
}

// Initialize se encarga de inyectar todas las dependencias usando Wire.
func Initialize() (*Dependencies, error) {
	wire.Build(
		// Proveedores bootstrap
		ProvideConfigLoader,
		ProvideGinServer,
		ProvideGormRepository,
		ProvideMongoDbRepository,
		ProvidePostgresRepository,
		ProvideJwtMiddleware,
		ProvideMiddlewares,
		ProvideRedisCache,
		ProvideJwtService,
		ProvideHttpClient,
		ProvideSmtpService,
		ProvideRabbitProducer,
		ProvideCassandraRepository,
		ProvideWebSocketUpgrader,

		// Person
		ProvidePersonRepository,
		ProvidePersonUseCases,
		ProvidePersonHandler,

		// Group
		ProvideGroupRepository,
		ProvideGroupUseCases,
		ProvideGroupHandler,

		// Event
		ProvideEventRepository,
		ProvideEventUseCases,
		ProvideEventHandler,

		// User
		ProvideUserRepository,
		ProvideUserUseCases,
		ProvideUserHandler,

		// Assessment
		ProvideAssessmentRepository,
		ProvideAssessmentUseCases,
		ProvideAssessmentHandler,

		// Candidate
		ProvideCandidateRepository,
		ProvideCandidateUseCases,
		ProvideCandidateHandler,

		// Browser Events
		ProvideBrowserEventsRepository,
		ProvideBrowserEventsUseCases,
		ProvideBrowserEventsWebsocket,
		ProvideBrowserEventsHandler,

		// Notification
		ProvideNotificationSmtpService,
		ProvideNotificationUseCases,
		ProvideNotificationHandler,

		// Authe
		ProvideAutheCache,
		ProvideAutheHttpClient,
		ProvideAutheJwtService,
		ProvideAutheUseCases,
		ProvideAutheHandler,

		// Tweet
		ProvideTweetBroker,
		ProvideTweetCache,
		ProvideTweetRepository,
		ProvideTweetUseCases,
		ProvideTweetHandler,

		// Item
		ProvideItemRepository,
		ProvideItemUseCases,
		ProvideItemHandler,

		// Category
		ProvideCategoryRepository,
		ProvideCategoryUseCases,
		ProvideCategoryHandler,

		// MacroCategory
		ProvideMacroCategoryRepository,
		ProvideMacroCategoryUseCases,
		ProvideMacroCategoryHandler,

		// Supplier
		ProvideSupplierRepository,
		ProvideSupplierUseCases,
		ProvideSupplierHandler,

		wire.Struct(new(Dependencies), "*"),
	)
	return &Dependencies{}, nil
}
