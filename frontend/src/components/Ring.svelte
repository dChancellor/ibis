<script lang="ts">
  type RingType = 'timed' | 'set';
  export let startingColor: string;
  export let endColor: string;
  export let innerDiameter: number;
  export let outerDiameter: number;

  export let type: RingType;
  export let animationLength: number;
  export let percentFilled: number = null;

  let firstHalfDegreesRotated: number = 0;
  let secondHalfDegreesRotated: number = 0;

  if (percentFilled > 50) {
    firstHalfDegreesRotated = 225 - 180;
    secondHalfDegreesRotated = 225 + (percentFilled - 50) * 3.6;
  } else {
    firstHalfDegreesRotated = 225 + percentFilled * 3.6;
    secondHalfDegreesRotated = 225;
  }
</script>

{#if type === 'timed'}
  <div
    class={'timer'}
    style={`--outside-circle-width: ${outerDiameter}px; 
    --inside-circle-width: ${innerDiameter}px; 
    --animation-length: ${animationLength}s;
    --hand-color: conic-gradient(${startingColor}, ${endColor})
    `}
  >
    <div class="hand"><span></span></div>
    <div class="hand"><span></span></div>
  </div>
{:else if type === 'set'}
  <div
    class={'timer'}
    style={`--outside-circle-width: ${outerDiameter}px; 
    --inside-circle-width: ${innerDiameter}px; 
    --hand-color: conic-gradient(${startingColor}, ${endColor})
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
    background: var(--hand-color);
    border-radius: 50%;
    height: var(--outside-circle-width);
    position: absolute;
    aspect-ratio: 1/1;
    transform: translate(-50%, -50%);
    top: 50%;
    left: 50%;
  }

  .timer:before {
    content: '';
    border-radius: 50%;
    height: calc(var(--outside-circle-width) - 5px);
    position: absolute;
    aspect-ratio: 1/1;
    transform: translate(-50%, -50%);
    top: 50%;
    left: 50%;
    box-shadow: 0px 0px 0px 10px #111;
  }

  .timer:after {
    background: var(--background-color);
    border-radius: 50%;
    content: '';
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
    border: solid rgba(0, 0, 0);
    border-width: calc(var(--outside-circle-width) / 2);
    border-top-color: transparent;
    border-right-color: transparent;
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
