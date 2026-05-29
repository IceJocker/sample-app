package main
import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)
func main() {
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)
	http.ListenAndServe(":8080", nil)
}
func blueHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.
gcloud container clusters get-credentials hello-cluster --zone us-west1-a
# Expose the services (This works now that kubectl is connected)
kubectl expose deployment development-deployment --name=dev-deployment-service --type=LoadBalancer --port=8080 --target-port=8080 --namespace=dev || true
kubectl expose deployment production-deployment --name=prod-deployment-service --type=LoadBalancer --port=8080 --target-port=8080 --namespace=prod || true

# Deploy Dev v2.0
git checkout dev
cat << 'EOF' > main.go
package main
import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"
)
func main() {
	http.HandleFunc("/blue", blueHandler)
	http.HandleFunc("/red", redHandler)
	http.ListenAndServe(":8080", nil)
}
func blueHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
func redHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
