import { writable } from "svelte/store";

export const todos = writable(
    [
        // {
        //     id: "tdl0",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl1",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl2",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl3",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl4",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl5",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl6",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
        // {
        //     id: "tdl7",
        //     title: "Lorem Ipsum",
        //     items: [
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: true,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //         {
        //             completed: false,
        //             name: "Do this, do that",
        //         },
        //     ],
        // },
    ]
)