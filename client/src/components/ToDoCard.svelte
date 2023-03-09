<script>
    let todo = {
        "title": "Lorem Ipsum",
        "items": [
            {
                "completed": false,
                "name": "Do this, do that"
            },
            {
                "completed": true,
                "name": "Do this, do that"
            },
            {
                "completed": false,
                "name": "Do this, do that"
            },
            {
                "completed": true,
                "name": "Do this, do that"
            },
            {
                "completed": false,
                "name": "Do this, do that"
            },
            {
                "completed": false,
                "name": "Do this, do that"
            },
            {
                "completed": false,
                "name": "Do this, do that"
            },
            {
                "completed": false,
                "name": "Do this, do that"
            }
        ]
    }


</script>

<main>
    <div class="card">
        <h2>{todo.title}</h2>
        <div class="items">
            {#each todo.items.slice(0, 5) as item}
                <div id="1" class="item">
                    <div class="checkbox done_{item.completed}"></div>
                    <div class="name done_{item.completed}">{item.name}</div>
                </div>
            {/each}
            {#if todo.items.length > 5}
                <div class="extra">. . .</div>
            {/if}
        </div>
    </div>
</main>

<style>
    .card {
        background-color: white;
        border: 2px solid black;
        border-radius: 10px;
        padding: 10% 7%;
        box-shadow: 2px 2px 0 0px black;
        min-height: 175px;
    }
    .card:hover {
        background-color: rgb(255, 255, 97);
        box-shadow: 7px 7px 0 0px black;
        cursor: pointer;
        transition: box-shadow 0.2s ease-in-out;
    }

    .items {
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
        width: 10px;
        height: 10px;
        border-radius: 3px;
        border: 2px solid black;
    }

    .extra {
        font-size: large;
        font-weight: bold;
    }

    .name.done_true {
        text-decoration: line-through;
    }
    .checkbox.done_true {
        background-color: black;
    }
</style>