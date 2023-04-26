package imageprocessing

import (
	"image"
)

// Структура, яка представляє зображення
type Image struct {
	data image.Image
}

// Функція, яка змінює розмір зображення
func (img *Image) Resize(width, height int) {
	// Реалізація зміни розміру зображення
}

// Функція, яка змінює яскравість зображення
func (img *Image) AdjustBrightness(brightness float64) {
	// Реалізація зміни яскравості зображення
}

// Функція, яка застосовує ефект "чорно-білий"
func (img *Image) ApplyBlackAndWhiteEffect() {
	// Реалізація ефекту "чорно-білий"
}

// Функція, яка застосовує ефект "розмиття"
func (img *Image) ApplyBlurEffect(radius float64) {
	// Реалізація ефекту "розмиття"
}

// Функція, яка застосовує ефект "відблисків"
func (img *Image) ApplyGlareEffect(strength float64) {
	// Реалізація ефекту "відблисків"
}

// Функція, яка застосовує ефект "відблисків"
func (img *Image) Do() error {
	// Реалізація ефекту "відблисків"
	return nil
}

// Всі методи працюють з об'єктом Image та надають різні методи зміни зображення - зміна розміру, яскравості, ефекти "чорно-білий",
//"розмиття" та "відблисків". Всі ці методи мають спільну властивість - вони працюють з об'єктом Image та змінюють його.
//Така організація коду дозволяє зберігати логіку всієї обробки зображень в одному місці та легко додавати нові методи для обробки зображень,
//які також працюватимуть з об'єктом Image та будуть зберігати високу когезію з іншими методами в класі Image.
