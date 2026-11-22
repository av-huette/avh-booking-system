package ws

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/av-huette/avh-booking-system/internal/models"
	"github.com/av-huette/avh-booking-system/internal/repo"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

// TestMain runs the tests from backend/, where the validator expects ../schemas, as the server does.
func TestMain(m *testing.M) {
	if err := os.Chdir("../.."); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

// resultRows returns one row per table, shaped as the database returns it (ids and timestamps set).
func resultRows() map[repo.TableName]any {
	createdAt := pgtype.Timestamp{Time: time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC), Valid: true}

	account := models.CreateAccount("Ada", "", "Lovelace", "ada@example.org", "", 1250, 5000, 1)
	account.ID, account.Enabled, account.CreatedAt = 1, true, createdAt
	category := models.CreateCategory("Members", "user", "account")
	category.ID, category.Enabled = 1, true
	location := models.CreateLocation("Bar")
	location.ID = 1
	order := models.CreateOrder(1)
	order.ID, order.CreatedAt = 1, createdAt
	product := models.CreateProduct("Beer", 250, 1, 1, 500, 1, 1)
	product.ID, product.CreatedAt = 1, createdAt
	parentID := 1
	productGroup := models.CreateProductGroup("Lager", &parentID)
	productGroup.ID = 2
	visibility := models.CreateProductVisibility(1, 1, 1)
	visibility.ID = 1
	rights := models.CreateRights(1, "booking", true)
	rights.ID = 1
	role := models.CreateRole("Admin")
	role.ID = 1
	sewobe := models.CreateServiceSewobe("key", "https://example.org")
	sewobe.ID = 1
	unit := models.CreateUnit("ml")
	unit.ID = 1
	user := models.CreateUser("admin", 1, "secret")
	user.ID = 1
	vat := models.CreateVat(19)
	vat.ID = 1
	setting := models.CreateSetting("compTitle", "AVH")

	return map[repo.TableName]any{
		repo.TableAccount:           account,
		repo.TableAccountOption:     models.CreateAccountOption(1, "theme", "dark"),
		repo.TableCategory:          category,
		repo.TableFavorites:         models.CreateFavorite(1, 1, 3),
		repo.TableLocation:          location,
		repo.TableOrder:             order,
		repo.TableProduct:           product,
		repo.TableProductGroup:      productGroup,
		repo.TableProductOrder:      models.CreateProductOrder(1, 1, 2, 250),
		repo.TableProductVisibility: visibility,
		repo.TableRights:            rights,
		repo.TableRole:              role,
		repo.TableServiceLink:       models.CreateServiceLink(42, 1, 1),
		repo.TableServiceSewobe:     sewobe,
		repo.TableSettingsEmail:     setting,
		repo.TableSettingsFrontend:  setting,
		repo.TableSettingsPayment:   setting,
		repo.TableUnit:              unit,
		repo.TableUser:              user,
		repo.TableUserOption:        models.CreateUserOption(1, "language", "de"),
		repo.TableVat:               vat,
	}
}

// TestResultSchemasMatchModels checks every table's model against its result schema,
// as one record and as a list, through the helpers the server sends results with.
func TestResultSchemasMatchModels(t *testing.T) {
	s := &Service{validator: newMessageValidator()}
	for table, row := range resultRows() {
		t.Run(string(table), func(t *testing.T) {
			raw, err := json.Marshal(row)
			require.NoError(t, err)

			_, wsErr := s.marshalBroadcastQueryResult(table, raw)
			require.Nil(t, wsErr, "one record")

			list := []json.RawMessage{raw}
			_, wsErr = s.marshalBroadcastQueryResultList(table, &list)
			require.Nil(t, wsErr, "list")
		})
	}
}

// TestResultSchemaFollowsTable checks that a row is validated against its table's schema, not any schema.
func TestResultSchemaFollowsTable(t *testing.T) {
	s := &Service{validator: newMessageValidator()}
	raw, err := json.Marshal(resultRows()[repo.TableProduct])
	require.NoError(t, err)

	_, wsErr := s.marshalBroadcastQueryResult(repo.TableUnit, raw)
	require.NotNil(t, wsErr, "a product sent as a unit is rejected")

	list := []json.RawMessage{raw}
	_, wsErr = s.marshalBroadcastQueryResultList(repo.TableUnit, &list)
	require.NotNil(t, wsErr, "a list of products sent as units is rejected")
}
