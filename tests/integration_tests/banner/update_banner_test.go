package banner

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ilyushkaaa/banner-service/tests/test_json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateBanner(t *testing.T) {
	tf := newBannerTestFixtures(t)
	defer tf.Close(t)

	t.Run("error duplicate feature + tag", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPatch, "/banner/1", strings.NewReader(test_json.BannerUpdateDuplicate))
		request = mux.SetURLVars(request, map[string]string{"id": "1"})
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.UpdateBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		body, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"such pair of feature and tag already exists"}`, string(body))
	})

	t.Run("error banner not found", func(t *testing.T) {
		setUp(t, tf.db, tableNames)
		fillDataBase(t, tf.db)
		request := httptest.NewRequest(http.MethodPatch, "/banner/10", strings.NewReader(test_json.BannerUpdateOk))
		request = mux.SetURLVars(request, map[string]string{"id": "10"})
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.UpdateBanner(w, r)
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
		request := httptest.NewRequest(http.MethodPatch, "/banner/1", strings.NewReader(test_json.BannerUpdateOk))
		request = mux.SetURLVars(request, map[string]string{"id": "1"})
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("Token", "admin_token")
		respWriter := httptest.NewRecorder()

		tf.mw.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tf.del.UpdateBanner(w, r)
		})).ServeHTTP(respWriter, request)
		resp := respWriter.Result()
		_, err := io.ReadAll(resp.Body)

		require.NoError(t, err)
		defer resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
	})
}
