import {
    Draggable,
    Sortable,
    Droppable,
    Swappable,
} from 'https://cdn.jsdelivr.net/npm/@shopify/draggable/build/esm/index.mjs';

const containers = document.querySelectorAll('.block')
const droppable = new Droppable(containers, {
    draggable: '.draggable',
    droppable: '.droppable'
});
droppable.on('drag:start', () => console.log('drag:start'));
droppable.on('droppable:over', () => console.log('droppable:over'));
droppable.on('droppable:out', () => console.log('droppable:out'));
