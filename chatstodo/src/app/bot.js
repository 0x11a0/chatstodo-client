import { ListEntry, LinkButton, MobileListEntry, MobileLinkButton } from "./components";

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



export function MobileBots() {

    return (
        <div className="
            flex-1
            flex-col
            justify-start items-center
            p-4 w-full
        ">
            <h3 className="
                text-dark-font-primary
                font-primary font-bold
                text-m my-2 text-center

            ">Chats your bot can see</h3>

            <MobileListEntry>
                <span>Telegram Chat 1</span>
                <MobileLinkButton title="Active" isActivated />
            </MobileListEntry>

            <MobileListEntry>
                <span>Telegram Chat 2</span>
                <MobileLinkButton title="Active" isActivated />
            </MobileListEntry>

            <MobileListEntry>
                <span>Telegram Chat 3</span>
                <MobileLinkButton title="Activate" />
            </MobileListEntry>
        </div>

    );

}



