<script lang="ts">
    export let text: string
    export let onClick: () => void
</script>

<button on:click={onClick} on:keypress={onClick}>{text}</button>

<style>
    button {
        all: unset;            /* Resets all button-specific styles */
        box-sizing: border-box; /* Ensures padding and border are included in width and height */
        display: inline-block;  /* Allows padding and margins to be applied as intended */
        cursor: pointer;        /* Makes sure the button has the pointer cursor */
        user-select: none;      /* Prevents text selection on click */
        background-color: transparent; /* Removes any background */
        border: none;           /* Removes the default border */
        padding: 0;             /* Removes default padding */
        margin: 0;              /* Removes default margin */
        font-family: inherit;   /* Inherits the font style from the parent */
        font-size: inherit;     /* Inherits the font size from the parent */
        line-height: inherit;   /* Inherits the line height from the parent */
        text-align: center;     /* Centers text horizontally */
        -webkit-appearance: none; /* Removes additional styling on WebKit browsers */
    }

    button:focus {
        outline: none; /* Removes focus outline when button is clicked */
    }
</style>