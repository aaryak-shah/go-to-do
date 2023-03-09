<script>
    import { prevent_default } from "svelte/internal";
    import { login, register } from "../requests/auth";
    import { notif } from "../stores/notif";
    import { user } from "../stores/user";

    let action = "Login";
    const toggle = () => {
        action === "Register" ? (action = "Login") : (action = "Register");
    };

    let userauth = {
        Email: "",
        Password: "",
    };

    let emailValidationText = "";
    let passValidationText = "";

    const emailValidator = () => {
        if (userauth.Email.length === 0) {
            emailValidationText = "Please enter an email";
            return false;
        } else if (
            userauth.Email.match(/^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/g) === null
        ) {
            emailValidationText = "Not a valid email";
            return false;
        } else {
            emailValidationText = "";
            return true;
        }
    };

    const passValidator = () => {
        if (userauth.Password.length === 0) {
            passValidationText = "Please enter a password";
            return false;
        } else if (
            userauth.Password.length < 6 ||
            userauth.Password.includes(" ")
        ) {
            passValidationText =
                "Not a valid password. Password should be at least 6 characters and contain no spaces";
            return false;
        } else {
            passValidationText = "";
            return true;
        }
    };

    const submitAction = async () => {
        if (emailValidator() && passValidator()) {
            try {
                let result =
                    action === "Register"
                        ? await register(userauth)
                        : await login(userauth);
                console.log(result);
                if (action === "Register") action = "Login";
                else {
                    user.set({ email: result.email });
                    console.log("user", $user.email);
                }
            } catch (err) {
                notif.set({
                    context: "error",
                    content: err.response.data.error,
                });
                console.error(err.response.data.error);
            }
        }
    };
</script>

<main>
    <form:prevent_default>
        <div class="card">
            <h1>{action}</h1>
            <div class="toggle" on:click={toggle} on:keypress={undefined}>
                {action === "Register" ? "Login" : "Register"} Instead
            </div>
            <div
                class="validation email-val {emailValidationText === ''
                    ? 'hide'
                    : ''}"
            >
                {emailValidationText}
            </div>
            <input
                bind:value={userauth.Email}
                type="email"
                name="email"
                id="email"
                placeholder="Your Email"
                on:input={emailValidator}
            />
            <div
                class="validation pass-val {passValidationText === ''
                    ? 'hide'
                    : ''}"
            >
                {passValidationText}
            </div>
            <input
                bind:value={userauth.Password}
                type="password"
                name="password"
                id="password"
                placeholder="Your Password"
                on:input={passValidator}
            />
            <input type="submit" on:click={submitAction} />
        </div>
    </form:prevent_default>
</main>

<style>
    .card {
        background-color: white;
        border: 2px solid black;
        border-radius: 15px;
        padding: 10%;
        box-shadow: 10px 10px 0 0px black;
        min-height: 175px;
        display: flex;
        flex-direction: column;
        align-items: start;
        gap: 15px;
    }

    .toggle {
        border: none;
        background-color: transparent;
        padding: 0;
        text-decoration: underline;
        font-size: 16px;
        cursor: pointer;
    }

    input {
        width: 100%;
        box-sizing: border-box;
        border: 2px solid black;
        border-radius: 10px;
        padding: 16px 18px;
        font-size: 15px;
        font-family: "Montserrat", sans-serif;
    }
    input::placeholder {
        font-size: 15px;
    }

    input[type="submit"] {
        cursor: pointer;
        font-weight: 900;
        text-transform: uppercase;
        background-color: black;
        color: white;
    }
    input[type="submit"]:hover {
        background-color: rgb(255, 255, 97);
        color: black;
        transition: all 0.2s ease-in-out;
    }

    .validation {
        color: red;
        font-weight: bold;
    }

    .hide {
        display: none;
    }
</style>
