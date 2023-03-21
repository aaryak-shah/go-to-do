import { writable } from "svelte/store";

export const notif = writable({
    "context": "",
    "content": ""
})