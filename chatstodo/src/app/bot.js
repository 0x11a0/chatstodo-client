import { ListEntry, LinkButton } from "./components";

export function BotPanel() {

    return (
        <div className="
            flex-1
            flex flex-col
            justify-start items-center
            p-4 w-full
        ">
            <h3 className="
                text-dark-font-primary
                font-primary font-bold
                text-m

            ">Chats your bot can see</h3>

            <ListEntry>
                <span>Telegram Chat 1</span>
                <LinkButton title="Active" isActivated />
            </ListEntry>

            <ListEntry>
                <span>Telegram Chat 2</span>
                <LinkButton title="Active" isActivated />
            </ListEntry>

            <ListEntry>
                <span>Telegram Chat 3</span>
                <LinkButton title="Activate" />
            </ListEntry>
        </div>

    );

}



