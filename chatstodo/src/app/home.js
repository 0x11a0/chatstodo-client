import { Card, CardEntry, CardEntrySubtitle } from "./card";

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

function TodoCard({ date }) {
    //const [todoList, setTodoList] = useState([]);
    const todoList = [
        {
            id: 1,
            title: "Todo 1",
            date: "11/02/2024",
            text: "Do laundry",
            group: "urgent",
            status: "completed"
        },
        {
            id: 2,
            title: "Todo 2",
            date: "11/02/2024",
            text: "Do this",
            group: "urgent",
            status: "incomplete"
        },
        {
            id: 3,
            title: "Todo 3",
            date: "12/02/2024",
            text: "Do that",
            group: "other",
            status: "incomplete"
        },
    ];

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

function EventsCard({ date }) {
    //const [todoList, setTodoList] = useState([]);
    const eventsList = [
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
    ];

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


function SummaryCard({ date }) {
    //const [todoList, setTodoList] = useState([]);
    const summaryList = [
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
    ];

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
