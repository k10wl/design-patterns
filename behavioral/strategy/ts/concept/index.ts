interface Strategy {
  calculateProfit(income: number):void
}

class StrategyContext {
  protected strategy: Strategy
  income: number

  constructor(strategy: Strategy) {
    this.strategy = strategy
    this.income = 420
  }

  setStrategy(strategy: Strategy): void {
    this.strategy =strategy
  }

  action() {
this.strategy.calculateProfit(this.income)
  }
}

class TransparentStrategy implements Strategy {
  calculateProfit(income: number): void {
      console.log('[clear]: checking total income')
      console.log('[clear]: subtracting payments')
      console.log(`[clear]: total profit is $${income}`)
  }
}

class RiggedStrategy implements Strategy {
  calculateProfit(income: number): void {
      const needToReport = 69.69
      const stolen = income - needToReport
      console.log(`[rigged]: total profit is $${needToReport} *wink*`)
      console.log('[rigged]: ...')
      console.log(`[rigged]: ha-ha I stole $${stolen}`)
  }
}

const transparent = new TransparentStrategy()
const strategyContext = new StrategyContext(transparent)
strategyContext.action()
const rigged = new RiggedStrategy()
strategyContext.setStrategy(rigged)
strategyContext.action()
