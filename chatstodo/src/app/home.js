import { FaSmileBeam } from "react-icons/fa";
import { Card, CardEntry, CardEntrySubtitle, MobileCard, MobileCardEntry } from "./card";
import { useState } from "react";

export function HomePanel() {
    const today = new Date();
    const day = String(today.getDate()).padStart(2, "0");
    const month = String(today.getMonth() + 1).padStart(2, "0");
    const year = today.getFullYear();
    const todayString = day + "/" + month + "/" + year;
    console.log(todayString);

    return (
        <div className="
            flex flex-wrap
            justify-start items-start
            p-2 w-full
        "

        >
            <TodoCard date={todayString} />
            <EventsCard date={todayString} />
            <SummaryCard date={todayString} />
        </div>

    );

}

const todoList = [
    {
        id: 1,
        title: "Fill indemnity form",
        date: "14/02/2024",
        text: "Needs to completed ASAP",
        group: "urgent",
        status: "incomplete"
    },
    {
        id: 2,
        title: "Research places with gin tonic",
        date: "11/02/2024",
        text: "",
        group: "other",
        status: "incomplete"
    },
    {
        id: 3,
        title: "Make a reservation for LAVO on March 18, 3pm",
        date: "11/02/2024",
        text: "",
        group: "other",
        status: "incomplete"
    },
    {
        id: 4,
        title: "Bring cake for the celebration",
        date: "11/02/2024",
        text: "",
        group: "other",
        status: "incomplete"
    },
];

function TodoCard({ date }) {
    //const [todoList, setTodoList] = useState([]);

    let urgentList = [];
    let otherList = [];
    todoList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    todoList.sort((todo1, todo2) => todo1.time - todo2.time);

    todoList.forEach((todo, index) => {
        if (todo.group === "urgent") {
            urgentList.push(<CardEntry key={index} todo={todo} date={date} />);
        } else {
            otherList.push(<CardEntry key={index} todo={todo} date={date} />);
        }
    });

    return (
        <Card title="Todo List" >
            <CardEntrySubtitle title="Urgent" />
            {urgentList}
            <CardEntrySubtitle title="Others" />
            {otherList}
        </Card>
    );
}

const eventsList = [
    /*
    {
        id: 1,
        title: "Event 1",
        date: "11/02/2024",
        text: "Laundry",
        group: "urgent",
        status: "completed"
    },
    {
        id: 2,
        title: "Booya",
        date: "11/02/2024",
        text: "Abc123",
        group: "urgent",
        status: "incomplete"
    },
    {
        id: 3,
        title: "Hello",
        date: "12/02/2024",
        text: "World",
        group: "other",
        status: "incomplete"
    },
    */
    {
        id: 1,
        title: "IDP outing at KSTAR",
        date: "17/03/2024",
        timeStr: "2100",
        text: "Description",
        type: "event"
    },
    {
        id: 2,
        title: "Antoine outing at LAVO",
        date: "18/03/2024",
        timeStr: "1500",
        text: "Description",
        type: "event"
    },
];
function EventsCard({ date }) {

    let cardList = [];
    eventsList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    eventsList.sort((todo1, todo2) => todo1.time - todo2.time);

    eventsList.forEach((todo, index) => {
        cardList.push(<CardEntry key={index} todo={todo} date={date} />);
    });

    return (
        <Card title="Events" >
            {cardList}
        </Card>
    );
}

const summaryList = [
    /*
    {
        id: 1,
        title: "Summary 1",
        date: "11/02/2024",
        mainText: [
            "CS202 Project: Planning meeting soon. Confirm your availability.",
            "CS460 Research: Discussing findings, meeting to finalize the research proposal.",
            "CS206 Coding: Tuesday session, review code snippets before.",
            "CS205 Exam Prep: Collaborate for resources, stay updated on deadlines.",
            "CCA - Samba Masala: Open house gig pre, practice on Tuesday and Thursday.",
        ],
        importantNotes: [
            "Prioritise CS202 meeting and CS420 research proposal.",
            "Review CS206 code before Tuesday.",
            "Collaborate on CS205 exam prep.",
            "Attend Samba Masala practices for upcoming open house gig.",
        ],
    },
    */
    {
        id: 1,
        title: "Validation Group 1",
        date: "11/02/2024",
        mainText: [
            `The group chat discussion revolved around planning two outings - 
            one for the IDP group and another to celebrate with a friend named Antoine.`,
            "A supper outing at Swee Choon was also mentioned.",
            `The conversation included logistical questions about fetching, sleeping
            schedules, and venue preferences.`,
            `Tasks were delegated for researching places and making reservations`,
            `Two specific events were finalised with dates and times.`

        ],
        importantNotes: [
            "Prioritise CS202 meeting and CS420 research proposal.",
            "Review CS206 code before Tuesday.",
            "Collaborate on CS205 exam prep.",
            "Attend Samba Masala practices for upcoming open house gig.",
        ],
    },
];

function SummaryCard({ date }) {

    let cardList = [];
    summaryList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    summaryList.sort((todo1, todo2) => todo1.time - todo2.time);

    summaryList.forEach((todo, index) => {
        cardList.push(<SummaryEntry key={index} todo={todo} date={date} />);
    });

    return (
        <Card title="Summary" type="summary">
            {cardList}
        </Card>
    );
}


function SummaryEntry({ todo, date }) {
    const dateTag = date === todo.date ?
        <p className="
            font-primary font-bold
            text-sm text-dark-font-primary
            bg-[#f8f8f8]/5 px-2
            inline-block
        "
        >Today</p> :
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
                font-primary font-bold
                text-lg text-dark-font-secondary
            ">{todo.title}</p>

            {dateTag}

            <ul className="list-disc px-4">
                {todo.mainText.map((point, index) =>

                (<li key={index}
                    className="
                        my-2
                        font-primary
                        text-m text-dark-font-secondary
                ">{point}</li>)

                )}

            </ul>
            {todo.importantNotes.length == 0 ? "" :
                <>
                    <p className="
                my-2
                font-primary font-bold
                text-m text-dark-font-secondary
            ">Important notes</p>

                    <ul className="list-disc px-4">
                        {todo.importantNotes.map((point, index) =>

                        (<li key={index}
                            className="
                        my-2
                        font-primary
                        text-m text-dark-font-secondary
                ">{point}</li>)

                        )}

                    </ul>
                </>
            }
        </div>
    );

}


export function MobileHome({ hasData }) {
    const today = new Date();
    const day = String(today.getDate()).padStart(2, "0");
    const month = String(today.getMonth() + 1).padStart(2, "0");
    const year = today.getFullYear();
    const todayString = day + "/" + month + "/" + year;
    console.log(todayString);

    return (
        <>
            {!hasData &&
                <>

                    <p className="
        text-dark-font-secondary font-primary text-center
        my-4
        "

                    >You have no upcoming tasks or events.<br />
                        Sync now to get the latest updates!
                    </p>
                    <FaSmileBeam className="
                text-4xl text-dark-font-secondary"
                    />
                </>
            }
            {hasData &&
                <>
                    <MobileTodoCard date={todayString} />
                    <MobileEventsCard date={todayString} />
                    <MobileSummaryCard date={todayString} />
                </>
            }
        </>
    );
}

function MobileTodoCard({ date }) {
    //const [todoList, setTodoList] = useState([]);

    let urgentList = [];
    let otherList = [];
    todoList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    todoList.sort((todo1, todo2) => todo1.time - todo2.time);

    todoList.forEach((todo, index) => {
        if (todo.group === "urgent") {
            urgentList.push(<MobileCardEntry key={index} todo={todo} date={date} />);
        } else {
            otherList.push(<MobileCardEntry key={index} todo={todo} date={date} />);
        }
    });

    return (

        <MobileCard title="Todo List" >
            <CardEntrySubtitle title="Urgent" />
            {urgentList}
            <CardEntrySubtitle title="Others" />
            {otherList}
        </MobileCard>

    );
}

function MobileEventsCard({ date }) {

    let cardList = [];
    eventsList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    eventsList.sort((todo1, todo2) => todo1.time - todo2.time);

    eventsList.forEach((todo, index) => {
        console.log(todo.time);
        cardList.push(<MobileCardEntry key={index} todo={todo} date={date} />);
    });

    return (
        <MobileCard title="Events" >
            {cardList}
        </MobileCard>
    );
}


function MobileSummaryCard({ date }) {

    let cardList = [];
    summaryList.forEach(todo =>
        todo.time = Date.parse(todo.date)
    );
    summaryList.sort((todo1, todo2) => todo1.time - todo2.time);

    summaryList.forEach((todo, index) => {
        cardList.push(<MobileSummaryEntry key={index} todo={todo} date={date} />);
    });

    return (
        <MobileCard title="Summary" type="summary">
            {cardList}
        </MobileCard>
    );
}


function MobileSummaryEntry({ todo, date }) {
    return (
        <div key={todo.id}
            className="
            border-b
            border-dark-border-primary
        "
        >
            <p className="
                my-2
                font-primary font-bold
                text-lg text-dark-font-secondary
            ">{todo.title}</p>


            <ul className="list-disc px-4">
                {todo.mainText.map((point, index) =>

                (<li key={index}
                    className="
                        my-2
                        font-primary
                        text-m text-dark-font-secondary
                ">{point}</li>)

                )}

            </ul>

            <p className="
                my-2
                font-primary font-bold
                text-m text-dark-font-secondary
            ">Important notes</p>

            <ul className="list-disc px-4">
                {todo.importantNotes.map((point, index) =>

                (<li key={index}
                    className="
                        my-2
                        font-primary
                        text-m text-dark-font-secondary
                ">{point}</li>)

                )}

            </ul>

        </div>
    );

}
