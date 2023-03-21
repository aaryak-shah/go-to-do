<script>
    import { todos } from "../../stores/todos";

    export let idx;
</script>

<main class="items">
    {#each $todos[idx].items.slice(0, 5) as item}
        <div id="1" class="item">
            <div class="checkbox done_{item.completed}" />
            <div class="name done_{item.completed}">{item.name}</div>
        </div>
    {/each}
    {#if $todos[idx].items.length > 5}
        <div class="extra">. . .</div>
    {/if}
</main>

<style>
    .items {
        width: 100%;
        overflow-x: auto;
        display: flex;
        flex-direction: column;
        gap: 10px;
    }

    .item {
        display: flex;
        flex-direction: row;
        justify-content: start;
        align-items: center;
        gap: 10px;
    }

    .checkbox {
        min-width: 10px;
        aspect-ratio: 1;
        border-radius: 3px;
        border: 2px solid black;
    }

    .extra {
        font-size: large;
        font-weight: bold;
    }

    .name {
        width: 100%;
        overflow-wrap: break-word;
    }
    .name.done_true {
        text-decoration: line-through;
    }
    .checkbox.done_true {
        background-color: black;
    }
</style>
