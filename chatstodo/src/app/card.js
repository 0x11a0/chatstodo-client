import { FaFileExport } from "react-icons/fa";

export function CardEntrySubtitle({ title }) {
    return (
        <h4 className="
            flex-1
            font-primary font-bold
            text-lg text-dark-font-primary
            my-2
        ">{title}</h4>

    );
}

export function Card({ title, type, children }) {
    const exportIcon = type === "summary" ?
        <div className="flex-1"></div> :
        <div className="
            flex-1 flex
            text-right justify-end
        ">
            <button>
                <FaFileExport className="
                text-right
                font-primary font-bold
                text-xl text-dark-font-primary
            "/>
            </button>
        </div>



    return (
        <div className="
            flex-1 flex-col justify-start items-start
            min-w-96
            bg-dark-background-secondary
            p-4 m-2 rounded-xl
            drop-shadow-xl
        "
        >
            <p className="
                text-center self-center
                font-primary font-bold
                text-lg
            ">...</p>
            <div className="
                flex flex-row justify-between items-center
                my-4
                ">

                <div className="flex-1"
                ></div>

                <h3 className="
                    flex-1
                    text-center
                    font-primary font-bold
                    text-xl text-dark-font-primary
                ">{title}</h3>

                {exportIcon}

            </div>
            {children}
        </div >
    );

}

export function CardEntry({ todo, date }) {
    const dateTag = date === todo.date ?
        <p className="
            font-primary font-bold
            text-sm text-dark-font-primary
            bg-[#f8f8f8]/5 px-2
            inline-block rounded
        "
        >Today&apos;s Task</p> :
        <p>{todo.date}</p>



    return (
        <div key={todo.id}
            className="
            border-b
            border-dark-border-primary
        "
        >
            <p className="
                my-2
                font-primary font-medium
                text-lg text-dark-font-secondary
            ">{todo.title}</p>

            {dateTag}

            <p className="
                my-2
                font-primary
                text-m text-dark-font-secondary
            ">{todo.text}</p>
        </div>
    );

}
