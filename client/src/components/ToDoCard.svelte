<script>
    import { selectedCard } from "../stores/selected";
    import CardItems from "./items/CardItems.svelte";
    import CardEditableItems from "./items/CardEditableItems.svelte";

    import { todos } from "../stores/todos";

    export let id;
    export let idx;

    const openAction = () => {
        if ($selectedCard !== id) {
            selectedCard.set(id);
        }
    };

    const closeAction = () => {
        if ($selectedCard === id) {
            selectedCard.set("");
        }
    };
</script>

<main {id}>
    <div class="backdrop expanded-{$selectedCard === id}" />
    {#if $selectedCard === id}
        <div class="close" on:click={closeAction} on:keypress={undefined}>
            &times;
        </div>
    {/if}
    <div
        class="card expanded-{$selectedCard === id}"
        on:click={openAction}
        on:keypress={undefined}
    >
        <h6>#{id}</h6>
        <h2>{$todos[idx].title}</h2>
        {#if $selectedCard === id}
            <CardEditableItems {idx} />
        {:else}
            <CardItems {idx} />
        {/if}
        <!-- {#each todo.items.slice(0, 5) as item}
                <div id="1" class="item">
                    <div class="checkbox done_{item.completed}" />
                    <div class="name done_{item.completed}">{item.name}</div>
                </div>
            {/each}
            {#if todo.items.length > 5}
                <div class="extra">. . .</div>
            {/if} -->
    </div>
</main>

<style>
    .backdrop {
        height: 0vh;
        width: 100vw;
        background-color: rgb(255, 255, 97);
        position: fixed;
        inset: 0;
        transition: height 0.3s ease;
    }
    .backdrop.expanded-true {
        height: 100vh;
    }

    .card {
        background-color: white;
        border: 2px solid black;
        border-radius: 10px;
        padding: 10% 7%;
        box-shadow: 2px 2px 0 0px black;
        min-height: 175px;

        transition: box-shadow 0.2s ease-in-out;
    }
    .card:hover {
        background-color: rgb(255, 255, 97);
        box-shadow: 7px 7px 0 0px black;
        cursor: pointer;
    }

    .expanded-true.card {
        position: fixed;
        top: 0;
        left: 0;
        bottom: 0;
        right: 0;
        margin: 75px 25%;
        box-shadow: 7px 7px 0 0px black;

        padding: 5%;
        overflow-y: auto;
    }
    .expanded-true.card:hover {
        background-color: white;
        cursor: default;
    }

    .close {
        width: 40px;
        height: 40px;
        border: 2px solid black;
        border-radius: 100px;
        background-color: white;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 36px;

        position: fixed;
        top: calc(75px - 21px);
        left: calc(25% - 21px);
        z-index: 2;
    }
    .close:hover {
        cursor: pointer;
        background-color: lightcoral;
    }

    h6 {
        color: rgba(0, 0, 0, 0.4);
        text-transform: uppercase;
        margin-bottom: 0;
        font-weight: 900;
    }
    h2 {
        margin-top: 0;
    }
</style>
