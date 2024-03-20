abstract class Creator {
  public abstract factoryMethod(): Product;

  public additionalOperation(): string {
    const product = this.factoryMethod();
    const result = product.operation();
    return `Product operation returned: ${result}`;
  }
}

class FooConcreteCreator extends Creator {
  public factoryMethod(): Product {
    return new FooProduct();
  }
}

class BarConcreteCreator extends Creator {
  public factoryMethod(): Product {
    return new BarProduct();
  }
}

interface Product {
  operation(): string;
}

class FooProduct implements Product {
  operation(): string {
    return "foo product opration";
  }
}

class BarProduct implements Product {
  operation(): string {
    return "bar product opration";
  }
}

function clientCode(creator: Creator) {
  console.log("Client: I only know that creator can call additionalOperation");
  console.log(creator.additionalOperation());
}

clientCode(new FooConcreteCreator());
clientCode(new BarConcreteCreator());
