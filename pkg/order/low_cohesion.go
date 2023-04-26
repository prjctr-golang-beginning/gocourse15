package order

// OrderProcessor відповідає за обробку замовлень
type OrderProcessor struct {
	validator  OrderValidator
	repository OrderRepository
	notifier   OrderNotifier
}

// NewOrderProcessor повертає новий об'єкт OrderProcessor з заданими залежностями
func NewOrderProcessor(validator OrderValidator, repository OrderRepository, notifier OrderNotifier) *OrderProcessor {
	return &OrderProcessor{validator, repository, notifier}
}

// ProcessOrder обробляє нове замовлення
func (op *OrderProcessor) ProcessOrder(o OrderIndependent) error {
	if err := op.validator.Validate(o); err != nil {
		return err
	}

	if err := op.repository.Save(o); err != nil {
		return err
	}

	if err := op.notifier.Notify(o); err != nil {
		return err
	}

	return nil
}

// OrderValidator валідує замовлення
type OrderValidator interface {
	Validate(o OrderIndependent) error
}

// OrderRepository зберігає замовлення
type OrderRepository interface {
	Save(o OrderIndependent) error
}

// OrderNotifier повідомляє про нове замовлення
type OrderNotifier interface {
	Notify(o OrderIndependent) error
}

// OrderIndependent відповідає за замовлення
type OrderIndependent struct {
	ID       int
	Customer string
	Items    []string
}

// Цей код демонструє low coupling між компонентами пакету order.
//Кожен компонент (OrderValidator, OrderRepository, OrderNotifier) використовується через інтерфейс,
//тому реалізація кожного компонента може бути змінена без впливу на інші компоненти. Крім того,
//у OrderProcessor є три залежності, але вони передаються через конструктор, що забезпечує зручну та гнучку конфігурацію залежностей.
