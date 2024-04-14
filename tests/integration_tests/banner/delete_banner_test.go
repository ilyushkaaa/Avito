package banner

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeleteBanner(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("banner not found", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodDelete, "/banner/10", nil)
		request = mux.SetURLVars(request, map[string]string{"id": "10"})
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.DeleteBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		_, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	})

	t.Run("ok", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodDelete, "/banner/1", nil)
		request = mux.SetURLVars(request, map[string]string{"id": "1"})
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.DeleteBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		_, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}
