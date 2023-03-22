<script>
    import { todos } from "../stores/todos";
    import ToDoCard from "./ToDoCard.svelte";

    const newList = () => {
        $todos = [
            { id: `tdl${$todos.length}`, items: [], title: "New ToDo" },
            ...$todos,
        ];
    };
</script>

<main>
    <div class="grid" id="todos">
        <div class="newlist" on:click={newList} on:keypress={undefined}>
            <div class="add">+</div>
            <div class="label">NEW TODO LIST</div>
        </div>
        {#each $todos as todo, idx (todo.id)}
            <ToDoCard id={todo.id} {idx} />
        {/each}
    </div>
</main>

<style>
    .grid {
        display: grid;
        grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
        grid-auto-rows: 1fr;
        gap: 15px;
    }

    .newlist {
        background-color: white;
        border: 2px solid black;
        border-radius: 10px;
        padding: 10% 7%;
        box-shadow: 5px 5px 0 0px black;
        min-height: 175px;

        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;

        transition: box-shadow 0.2s ease-in-out;
    }
    .newlist:hover {
        background-color: rgb(138, 222, 255);
        box-shadow: 7px 7px 0 0px black;
        cursor: pointer;
    }

    .add {
        font-size: 50px;
    }
</style>
