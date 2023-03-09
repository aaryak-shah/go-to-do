<script>
    import { notif } from "../stores/notif";

    export let context;
    export let content;

    const closeAction = () => {
        notif.set({
            context: "",
            content: "",
        });
    };
</script>

<main class="notif-pkg">
    <div class="notif">
        <div class="close" on:click={closeAction} on:keypress={undefined}>
            &times;
        </div>
        <div class="content ctx-{context}">{content}</div>
        <div class="timeout" />
    </div>
</main>

<style>
    .notif-pkg {
        position: fixed;
        width: 100%;
        top: -100px;
        animation-name: slideinout;
        animation-duration: 12s;
    }

    .close {
        width: 25px;
        height: 25px;
        border: 2px solid black;
        border-radius: 100px;
        background-color: white;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 24px;

        position: absolute;
        top: calc(0px - 12.5px);
        left: calc(25% - 12.5px);
    }
    .close:hover {
        cursor: pointer;
        background-color: rgb(138, 222, 255);
    }

    .notif {
        margin: 0px 25%;
        border: 2px solid black;
        border-radius: 10px;
        background-color: white;
        box-shadow: 7px 7px 0 0px black;
        display: flex;
        flex-direction: column;
        justify-content: start;
        overflow: hidden;
    }

    .content {
        padding: 20px;
        padding-bottom: 13px;
    }

    .ctx-error {
        color: red;
        font-weight: bold;
    }

    .timeout {
        width: 0%;
        height: 7px;
        background-color: lightcoral;
        animation-duration: 10s;
        animation-name: countdown;
        animation-iteration-count: 1;
    }

    @keyframes countdown {
        0% {
            width: 100%;
        }
        100% {
            width: 0%;
        }
    }

    @keyframes slideinout {
        0% {
            top: -100px;
        }
        8.3333% {
            top: 75px;
        }
        91.6663% {
            top: 75px;
        }
        100% {
            top: -100px;
        }
    }
</style>
