package imageprocessing

import (
	"image"
)

// ConvertToGray converts the input image to grayscale
func ConvertToGray(inputImg image.Image) image.Image {
	// ... apply grayscale filter ...
	return nil
}

// Resize resizes the input image to the given width and height
func Resize(inputImg image.Image, width, height int) image.Image {
	// ... apply resize algorithm ...
	return nil
}

//Якщо використовувати ці два методи в одній послідовності,
//наприклад, спочатку застосувавши фільтр до зображення та потім змінивши його розмір,
//можна отримати бажаний результат. Але, якщо спочатку змінити розмір зображення та потім
//застосувати фільтр, можна отримати зовсім інший результат. Це є прикладом низької зчепленості,
//оскільки два методи мають дуже мало спільного та можуть бути використані тільки в обмеженому контексті.
