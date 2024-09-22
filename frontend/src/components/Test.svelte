<div class="timer-group">
    <div class="timer hour">
      <div class="hand"><span></span></div>
      <div class="hand"><span></span></div>
    </div>
    <div class="timer minute">
      <div class="hand"><span></span></div>
      <div class="hand"><span></span></div>
    </div>
    <div class="timer second">
      <div class="hand"><span></span></div>
      <div class="hand"><span></span></div>
    </div>
  </div>


<style>
    .timer-group {
        height: 400px;
        margin: 0 auto;
        position: relative;
        aspect-ratio: 1/1;
    }

    .timer {
        background: rgba(0, 0, 0, .2);
        border-radius: 50%;
        height: var(--outside-circle-width);
        position: absolute;
        aspect-ratio: 1/1;
        transform: translate(-50%, -50%);
        top: 50%;
        left: 50%;
    }

    .timer:after {
        background: #111;
        border-radius: 50%;
        content: "";
        display: block;
        position: absolute;
        height: var(--inside-circle-width);
        aspect-ratio: 1/1;
        left: calc((var(--outside-circle-width) - var(--inside-circle-width)) / 2);
        top: calc((var(--outside-circle-width) - var(--inside-circle-width)) / 2);
    }

    .hand {
        float: left;
        height: 100%;
        overflow: hidden;
        position: relative;
        width: 50%;
    }

    .hand span {
        border: solid var(--hand-color);
        border-width: calc(var(--outside-circle-width) / 2);
        border-bottom-color: transparent;
        border-left-color: transparent;
        border-radius: 50%;
        display: block;
        height: 0;
        position: absolute;
        right: 0;
        top: 0;
        transform: rotate(225deg);
        width: 0;
        animation-duration: var(--animation-length);
        animation-iteration-count: infinite;
        animation-timing-function: linear;

    }
    .hand:first-child {
        transform: rotate(180deg);
    }
    
    .second {
        --outside-circle-width: 300px;
        --inside-circle-width: 290px;
        --hand-color: rgba(255, 255, 255, .15);
        --animation-length: 1s;
    }

    .minute {
        --outside-circle-width: 350px;
        --inside-circle-width: 300px;
        --hand-color: rgba(0, 255, 255, .4);
        --animation-length: 60s;
    }

    .hour {
        --outside-circle-width: 400px;
        --inside-circle-width: 350px;
        --hand-color: rgba(255, 0, 255, .4);
        --animation-length: 3600s;
    }

    .hand:first-child span {
        animation-name: spin1;  
    }

    .hand:last-child span {
        animation-name: spin2; 
    }

    @keyframes spin1 {

        50% {
            transform: rotate(225deg);
        }
        100% {
            transform: rotate(405deg);
        }
    }

    @keyframes spin2 {
        50% {
            transform: rotate(405deg);
        }
        100% {
            transform: rotate(405deg);
        }
    }
  </style>