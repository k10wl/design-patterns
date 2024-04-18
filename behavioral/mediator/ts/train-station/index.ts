interface IStation {
  notify(sender: object, data: "arrived" | "departed"): void;
}

class Train {
  name: string;
  #station: IStation;
  constructor(name: string, station: IStation) {
    this.name = name;
    this.#station = station;
  }
  arrive(): void {
    this.#station.notify(this, "arrived")
  }
  depart(): void {
    this.#station.notify(this, "departed")
  }
}

class Station {
  #queue: Train[]
  #onTrack: Train | null
  constructor() {
    this.#queue = []
    this.#onTrack = null
  }
  notify(sender: Train, data: "arrived" | "departed"): void {
    switch (data) {
      case "arrived":
        if (this.#onTrack) {
          console.log("+ waiting train:", sender.name)
          this.#queue.push(sender)
        } else {
          console.log("> train arrived on track:", sender.name)
          this.#onTrack = sender
        }
        break;
      case "departed":
        console.log("- departed:", sender.name)
        this.#onTrack = null
        if (this.#queue.length > 0) {
          this.#queue.shift()?.arrive()
        }
        break;
      default:
        throw new Error("Unhandled data in Station")
    }
  }
}

const station = new Station()
const trainA = new Train("train A", station)
const trainB = new Train("train B", station)
const trainC = new Train("train C", station)
const trainD = new Train("train D", station)
const trainE = new Train("train E", station)

trainA.arrive()
trainB.arrive()
trainC.arrive()
trainD.arrive()
trainE.arrive()

trainA.depart()
trainB.depart()
trainC.depart()
trainD.depart()

