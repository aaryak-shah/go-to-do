<script>
    import { todos } from "../../stores/todos";

    export let idx;

    const newItem = () => {
        $todos[idx].items = [
            { completed: false, name: "" },
            ...$todos[idx].items,
        ];
    };

    const deleteItem = (idx) => {};
</script>

<main class="items">
    <div class="delete-todo-btn">&times; Delete List</div>
    <div class="new-item-btn item" on:click={newItem} on:keypress={undefined}>
        <div class="new-icon">+</div>
        <div class="new-text">Add a new item</div>
    </div>
    {#each $todos[idx].items as item}
        <div id="1" class="item">
            <label class="checkbox-wrapper">
                <span class="checkbox-display" data-checked={item.completed}>
                    <input
                        class="checkbox"
                        type="checkbox"
                        bind:checked={item.completed}
                    />
                </span>
            </label>
            <input class="name done-{item.completed}" bind:value={item.name} />
            <div class="delete-item-btn">&times;</div>
        </div>
    {/each}
</main>

<style>
    .items {
        display: flex;
        flex-direction: column;
        gap: 16px;
        transition: all 0.1s;
    }

    .item {
        display: flex;
        flex-direction: row;
        justify-content: start;
        align-items: center;
        gap: 15px;
    }

    .checkbox-wrapper {
        position: relative;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        cursor: pointer;
    }

    .checkbox-display {
        background-color: white;
        border-radius: 3px;
        border: 2px solid black;
        cursor: pointer;
        width: 12px;
        height: 12px;

        transition: all 0.2s;
    }

    .checkbox-display[data-checked="true"] {
        background-color: black;
    }

    .checkbox {
        margin: 0;
        opacity: 0;
        width: 12px;
        height: 12px;
        cursor: pointer;
    }

    .extra {
        font-size: large;
        font-weight: bold;
    }

    .name {
        font-family: "Montserrat", sans-serif;
        font-size: 18px;
        border: none;
        outline: none;
        width: 100%;
        height: 100%;
        resize: none;
        transition: all 0.1s;
        overflow: hidden;
    }

    .name.done-true {
        text-decoration: line-through;
    }
    .checkbox.done-true {
        background-color: black;
    }

    .delete-item-btn {
        font-size: 26px;
        cursor: pointer;
        margin: 0;
        opacity: 0;
    }

    .item:hover .delete-item-btn {
        opacity: 100;
    }

    .item:hover .name,
    .name:focus,
    .name:active {
        border-bottom: 1px solid black;
    }

    .delete-todo-btn {
        border: 2px solid black;
        border-radius: 5px;
        background-color: lightcoral;
        width: max-content;
        padding: 10px;
        font-weight: bold;
        cursor: pointer;
        box-shadow: 2px 2px 0 0 black;
        transition: all 0.2s;

        margin-bottom: 20px;
    }
    .delete-todo-btn:hover {
        box-shadow: 5px 5px 0 0 black;
    }

    .new-item-btn {
        cursor: pointer;
    }
    .new-icon {
        width: 16px;
        font-size: 26px;
        color: gray;
    }
    .new-text {
        color: gray;
    }
</style>
