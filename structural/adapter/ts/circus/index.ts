interface ClientInterface {
  showOff(): void
}

class Adapter implements ClientInterface {
  adaptee: ClientInterface;

  constructor (concreteAdapter: ClientInterface) {
    this.adaptee = concreteAdapter
  }

  showOff(): void {
     this.adaptee.showOff() 
  }
}

class ClownAdapter implements ClientInterface {
  showOff(): void {
      new Clown().makeFun()
  }
}

class BearAdapter implements ClientInterface {
  showOff(): void {
      new Bear().rideBicycle()
  }
}

class ZebraAdapter implements ClientInterface {
  showOff(): void {
      new Zebra().run()
  }
}

class Clown {
  makeFun(): void {
    console.log('> clown throws water balloons')
  }
}

class Zebra {
  run(): void {
    console.log('> zebra runs around')
  }
}

class Bear {
  rideBicycle(): void {
    console.log('> bear rides a bicycle')
  }
}

function perform(actor: ClientInterface): void {
  actor.showOff()
}


console.log("> Adapter demo")
perform(new ClownAdapter())
perform(new BearAdapter())
perform(new ZebraAdapter())
