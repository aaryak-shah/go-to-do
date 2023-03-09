import { writable } from "svelte/store";

export const notif = writable({
    "context": "message",
    "content": "Some notification message"
})