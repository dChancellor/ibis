<script lang="ts">
  type RingType = "timed" | "set";
  export let color: string;
  export let innerDiameter: number;
  export let outerDiameter: number;

  export let type: RingType;
  export let animationLength: number;
  export let percentFilled: number = null;

  let firstHalfDegreesRotated: number = 0;
  let secondHalfDegreesRotated: number = 0;

  if (percentFilled > 50) {
    firstHalfDegreesRotated = 225 + 180;
    secondHalfDegreesRotated = 225 + (percentFilled - 50) * 3.6;
  } else {
    firstHalfDegreesRotated = 225 + percentFilled * 3.6;
    secondHalfDegreesRotated = 225;
  }
</script>

{#if type === "timed"}
  <div
    class={"timer"}
    style={`--outside-circle-width: ${outerDiameter}px; 
    --inside-circle-width: ${innerDiameter}px; 
    --hand-color: ${color}; 
    --animation-length: ${animationLength}s;`}
  >
    <div class="hand"><span></span></div>
    <div class="hand"><span></span></div>
  </div>
{:else if type === "set"}
  <div
    class={"timer"}
    style={`--outside-circle-width: ${outerDiameter}px; 
    --inside-circle-width: ${innerDiameter}px; 
    --hand-color: ${color};
    `}
  >
    <div class="hand">
      <span style={`transform: rotate(${secondHalfDegreesRotated}deg);`}></span>
    </div>
    <div class="hand">
      <span style={`transform: rotate(${firstHalfDegreesRotated}deg);`}></span>
    </div>
  </div>
{/if}

<style>
  .timer {
    background: rgba(0, 0, 0, 0.2);
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
    border: solid;
    /* border: solid var(--hand-color); */
    border-width: calc(var(--outside-circle-width) / 2);
    border-bottom-color: transparent;
    border-left-color: transparent;
    border-radius: 50%;
    display: block;
    position: absolute;
    right: 0;
    top: 0;
    transform: rotate(225deg);
    animation-duration: var(--animation-length);
    animation-iteration-count: infinite;
    animation-timing-function: linear;
  }
  .hand:first-child {
    transform: rotate(180deg);
  }

  .hand:first-child span {
    animation-name: spin1;
    border-top-color: radial-gradient(rgb(224, 12, 12), rgb(210, 35, 35)) padding-box, radial-gradient(to left, darkblue, rgb(33, 17, 40)) border-box;
  }

  .hand:last-child span {
    animation-name: spin2;
    background: radial-gradient(white, white) padding-box, radial-gradient(to right, darkblue, rgb(33, 17, 40)) border-box;
    /* border-top-color: linear-gradient(white, white) padding-box, linear-gradient(to left, darkblue, rgb(33, 17, 40)) border-box; */
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
