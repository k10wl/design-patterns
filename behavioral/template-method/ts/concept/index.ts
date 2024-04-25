class FinanceCalculator {
  income: number;

  constructor(income: number) {
    this.income = income;
  }

  calculate() {
    return this.getTaxes();
  }

  getTaxes(): number {
    return this.income * 0.05;
  }
}
class SmallBusinessCalculator extends FinanceCalculator {
  getTaxes(): number {
    return this.income * 0.03;
  }
}

class CorporationBusinessCalculator extends FinanceCalculator {
  getTaxes(): number {
    return this.income * 0.33;
  }
}

const profit = 100;
const smallBusinessCalculator = new SmallBusinessCalculator(profit);
const corpoBusinessCalculator = new CorporationBusinessCalculator(profit);

function client(calculator: FinanceCalculator): number {
  return calculator.calculate();
}

console.log(
  `> Small business taxes for ${profit} - ${client(smallBusinessCalculator)}`,
);
console.log(
  `> Corpo business taxes for ${profit} - ${client(corpoBusinessCalculator)}`,
);
