abstract class Invoker {
  #command: Command;
  constructor(command: Command) {
    this.#command = command;
  }
  invoke(): void {
    this.#command.execute();
  }
  setCommand(command: Command): void {
    this.#command = command;
  }
}

class Button extends Invoker {}
class KeyboardShortcut extends Invoker {}
class ContextMenuOption extends Invoker {}

class Receiver {
  operation(): void {
    console.log(">> receiver operation");
  }
}

interface Command {
  execute(): void;
}

class ReceiverCommand implements Command {
  #receiver: Receiver;
  constructor(receiver: Receiver) {
    this.#receiver = receiver;
  }
  execute(): void {
    this.#receiver.operation();
  }
  setReceiver(receiver: Receiver): void {
    this.#receiver = receiver;
  }
}

class LogCommand implements Command {
  param: string;
  constructor(param: string) {
    this.param = param;
  }
  execute(): void {
    console.log(">", this.param);
  }
}

const logCommand = new LogCommand("command as param");
const receiverCommand = new ReceiverCommand(new Receiver());

const button = new Button(logCommand);
const keyboardShortcut = new KeyboardShortcut(logCommand);
const contextMenuOption = new ContextMenuOption(logCommand);

button.invoke();
keyboardShortcut.invoke();
contextMenuOption.invoke();

button.setCommand(receiverCommand);
button.invoke();
