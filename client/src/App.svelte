<script>
    import { user } from "./stores/user";
    import { notif } from "./stores/notif";
    import Auth from "./components/Auth.svelte";
    import ToDos from "./components/ToDos.svelte";
    import { login, logout } from "./requests/auth";
    import { onMount } from "svelte/internal";
    import Notif from "./components/Notif.svelte";

    // onMount(async () => {
    //     try {
    //         let result = await login();
    //         user.set({ email: result.email });
    //     } catch (err) {
    //         console.error(err);
    //     }
    // });

    const logoutAction = async () => {
        try {
            await logout();
            user.set({ email: undefined });
        } catch (err) {
            console.error(err);
        }
    };
</script>

<div class="banner" />
<div class="notifarea {$notif.content === '' ? 'hide' : 'show'}">
    {#key $notif.content}
        <!-- {#if $notif.content !== ""} -->
        <Notif context={$notif.context} content={$notif.content} />
        <!-- {/if} -->
    {/key}
</div>
<main>
    <div class="header">
        <div class="left">
            <h1>GoToDo</h1>
            <div class="subtitle">Get Stuff Done</div>
        </div>
        {#if $user.email !== undefined}
            <div class="right">
                <h3 class="user">{$user.email}</h3>
                <div
                    class="logout"
                    on:click={logoutAction}
                    on:keypress={undefined}
                >
                    Log out
                </div>
            </div>
        {/if}
    </div>
    {#if $user.email === undefined}
        <Auth />
    {:else}
        <ToDos />
    {/if}
</main>

<style>
    main {
        padding: 0px 25% 160px 25%;
    }

    .header {
        margin-top: 160px;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    h1 {
        /* margin-top: 160px; */
        margin-bottom: 0px;
    }

    h3 {
        margin-bottom: 0;
    }

    .right {
        display: flex;
        flex-direction: column;
        align-items: end;
    }

    .logout {
        text-transform: uppercase;
        margin-bottom: 15px;
    }
    .logout:hover {
        text-decoration: underline;
        cursor: pointer;
    }

    .subtitle {
        text-transform: uppercase;
        /* font-weight: bold; */
        /* color: grey; */
        margin-bottom: 15px;
    }

    .banner {
        background-color: lightcoral;
        height: 480px;
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        z-index: -1;
        border-bottom: 2px solid black;
    }

    .hide {
        display: none;
    }
</style>
