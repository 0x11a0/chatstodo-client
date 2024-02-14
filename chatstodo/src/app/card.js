import { FaFileExport } from "react-icons/fa";
import { MdOutlineCheckBox, MdOutlineCheckBoxOutlineBlank } from "react-icons/md";
import { useState } from "react";

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
export function MobileCardEntrySubtitle({ title }) {
    return (
        <h4 className="
            text-start w-full
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
//<a href="https://calendar.google.com/" target="_blank">
//<a href="https://calendar.google.com/calendar?authuser=0">
export function MobileCard({ title, type, showCheckboxes, setShowCheckboxes, children, numTicks }) {
    const exportIcon = type === "summary" ?
        <div className="flex-1"></div> :
        <div className="
            flex-1 flex
            text-right justify-end
        ">
            <button onClick={() => setShowCheckboxes(!showCheckboxes)}

            >
                <FaFileExport className="
                text-right
                font-primary font-bold
                text-xl text-dark-font-primary
            "/>

            </button>
        </div>

    const [showExportButton, setShowExportButton] = useState(false);

    return (
        <div className="
            flex-1 flex flex-col justify-start items-center
            min-w-full
            bg-dark-background-secondary
            p-4 m-4 rounded-xl
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
                my-4 w-full
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
            {numTicks == 0 ? "" :
                <GoogleCalendarButton setShowCheckboxes={setShowCheckboxes}
                    setShowExportButton={setShowExportButton} />
            }
        </div >
    );

}

function GoogleCalendarButton({ setShowCheckboxes, setShowExportButton }) {



    return (
        <a href="https://calendar.google.com/calendar?authuser=0"
            className="
            flex-1
            font-primary font-bold self-center
            text-sm text-dark-font-primary
            bg-[#F8F8F8]/5px-2 mt-2
            rounded"
            onClick={() => {
                setShowCheckboxes(false);
                setShowExportButton(false);
            }}
        > Add to Google Calendar</a >
    );
}

export function MobileCardEntry({ todo, date, showCheckboxes, numTicks, setNumTicks }) {
    var dateTag;
    if (todo.type === "event") {
        const dateArr = todo.date.split("/");
        const eventDate = new Date(dateArr[1] + "-" + dateArr[0] + "-" + dateArr[2]);
        const dateFormatter = new Intl.DateTimeFormat('en-US', { day: 'numeric', month: 'long' });

        dateTag =
            <p className="
             font-primary font-bold
             text-sm text-dark-font-primary
             bg-[#F8F8F8]/5 px-2 mt-2
             inline-block rounded
                 "
            >{dateFormatter.format(eventDate)}, {todo.timeStr}</p>

    } else {
        dateTag = date === todo.date ?
            <p className="
        font-primary font-bold
        text-sm text-dark-font-primary
        bg-[#F8F8F8]/5 px-2 mb-2
        inline-block rounded
        "
            >Today&apos;s task</p> :
            <p></p>
    }

    const [isTicked, setIsTicked] = useState(false);


    return (
        <div key={todo.id}
            className="
            border-b w-full
            border-dark-border-primary
            pt-2
        "
        >
            <div className="
                flex flex-row justify-start
                items-center
    
        ">
                {!showCheckboxes ? "" :
                    (isTicked ?
                        <MdOutlineCheckBox className="
                            float-left 
                            text-lg text-dark-font-secondary
                            mr-2"
                            onClick={() => {
                                setIsTicked(false);
                                setNumTicks(numTicks - 1);
                            }}
                        />
                        :

                        <MdOutlineCheckBoxOutlineBlank className="
                            float-left 
                            text-lg text-dark-font-secondary
                            mr-2"
                            onClick={() => {
                                setIsTicked(true);
                                setNumTicks(numTicks + 1);
                            }}
                        />
                    )
                }
                <p className="
                font-primary font-medium
                text-m text-dark-font-secondary
            ">
                    {todo.title}</p>
            </div>
            {dateTag}

            {todo.type !== "event" ? "" :
                <p className="
                my-2
                font-primary
                text-m text-dark-font-secondary
            ">{todo.text}</p>
            }
        </div>
    );

}
