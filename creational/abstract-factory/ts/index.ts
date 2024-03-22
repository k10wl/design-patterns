// abstract factory
interface CommercialFactory {
  createOrder(id: number): Order;
  createPayment(): Payment
}

// abstract product1
interface Order {
  id: number;
  shipTo(destination: string): void;
  updateStorage(): void;
}

// abstract product2
interface Payment {
  savePayment(order: Order, amount: number): void
}

// concrete product1
class OnlineOrder implements Order {
  id: number;
  constructor(id: number) {
    this.id = id;
  }
  shipTo(destination: string): void {
    console.log(">>> shipping destination:", destination);
  }
  updateStorage(): void {
    console.log(">>> remove 1 from storage");
  }
}

// concrete product1
class PersonalOrder implements Order {
  id: number;
  constructor(id: number) {
    this.id = id;
  }
  shipTo(destination: string): void {
    console.log(">>> shipping destination:", destination);
  }
  updateStorage(): void {
    console.log(">>> remove 1 from local shop storage");
  }
}

// concrete product2
class OnlinePayment implements Payment {
  savePayment(order: Order, amount: number): void {
      console.log(`>>> payed online ${amount} social credits for order with id ${order.id}`)
  }
}

// concrete product2
class CashPayment implements Payment {
  savePayment(order: Order, amount: number): void {
      console.log(`>>> gave physically ${amount} social credits for order with id ${order.id}`)
  }
}

// concrete factory1
class CoffeeShop implements CommercialFactory {
  createOrder(id: number): Order {
    return new PersonalOrder(id);
  }
  createPayment(): Payment {
     return new CashPayment()
  }
}

// concrete factory2
class Marketpalce implements CommercialFactory {
  createOrder(id: number): Order {
    return new OnlineOrder(id);
  }
  createPayment() {
    return new OnlinePayment()
  }
}

function orderer(factory: CommercialFactory) {
  const order = factory.createOrder(performance.now());
  const payment = factory.createPayment();
  payment.savePayment(order, performance.now())
  order.updateStorage();
  order.shipTo("somewehere in ther world");
}

orderer(new CoffeeShop());
console.log("---");
orderer(new Marketpalce());
