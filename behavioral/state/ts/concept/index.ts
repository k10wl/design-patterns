class Context {
  #state: BaseState

  constructor(state: BaseState) {
    this.#state = state
    this.#state.setContext(this)
  }

  changeState(state: BaseState): void {
    this.#state = state
    this.#state.setContext(this)
  }

  public giveCoffee() {
    this.#state.giveCoffee()
  }

  public depositMoney() {
    this.#state.depositMoney()
  }
}

abstract class BaseState {
  context?: Context = undefined

  public setContext(context: Context): void {
    this.context = context
  }

  public abstract giveCoffee(): void

  public abstract depositMoney(): void
}

class NoMoneyState extends BaseState {
  public giveCoffee(): void {
    console.log("x ~ no money - no coffee!")
  }
  public depositMoney(): void {
    console.log("+ ~ deposited money")
    this.context?.changeState(new ReadyToServeState())
  }
}

class ReadyToServeState extends BaseState {
  public giveCoffee(): void {
    console.log("+ ~ coffee served")
    this.context?.changeState(new NoMoneyState())
  }
  public depositMoney(): void {
    console.log("x ~ return money, already deposited")
  }
}

const context = new Context(new NoMoneyState())

context.giveCoffee()
context.depositMoney()
context.giveCoffee()
context.giveCoffee()
context.giveCoffee()
context.giveCoffee()
context.giveCoffee()
context.depositMoney()
context.giveCoffee()
