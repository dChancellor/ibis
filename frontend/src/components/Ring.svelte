<script lang="ts">
  export let type: 'running' | 'static' = 'running';
  export let isRunning: boolean = true;
  export let diameter: number;
  export let duration: number;
  export let color: string;
  export let strokeWidth = 20;
  export let direction: 'up' | 'down';
  export let oneTime = false;
  export let progress = 0;

  let circumference = 2 * Math.PI * (diameter / 2);
  let id = crypto.randomUUID();

  let startTime: number;
  let currentProgress = 0;
  function animate(time: number) {
    if (!startTime) {
      startTime = time;
    }

    const elapsed = (time - startTime) / 1000;
    currentProgress = Math.min((elapsed / duration) * 100, 100);

    if (currentProgress > 99) {
      isRunning = false;
    } else if (currentProgress < progress) {
      requestAnimationFrame(animate);
    } else {
      isRunning = false;
    }
  }

  if (type === 'static') {
    requestAnimationFrame(animate);
  }
</script>

<div
  id="container"
  style={`
  --max-dash-length: ${circumference}; 
  --duration: ${duration}s; 
  --stroke-width:${strokeWidth}; 
  --color:${color}; 
  --progress: ${progress}; 
  --diameter: ${diameter + strokeWidth}px; 
  --animation-direction: ${direction === 'up' ? 'normal' : 'reverse'}; 
  --animation-iteration-count: ${oneTime ? 1 : 'infinite'}; 
  --animation-easing:${type === 'running' ? 'linear' : 'ease-in-out'}
  `}
>
  <svg
    width={diameter + strokeWidth}
    height={diameter + strokeWidth}
    viewBox="0 0 {diameter + strokeWidth} {diameter + strokeWidth}"
  >
    <defs>
      <mask id={`${id}-mask`}>
        <circle
          class="background-ring"
          cx={(diameter + strokeWidth) / 2}
          cy={(diameter + strokeWidth) / 2}
          r={diameter / 2}
        />
        <circle
          class="progress-ring"
          style={`animation-play-state:${isRunning ? 'running' : 'paused'}`}
          cx={(diameter + strokeWidth) / 2}
          cy={(diameter + strokeWidth) / 2}
          r={diameter / 2}
        />
        <circle
          style={`animation-play-state:${isRunning ? 'running' : 'paused'}`}
          cx={diameter + strokeWidth / 2}
          cy={(diameter + strokeWidth) / 2}
          r={strokeWidth / 2}
          class="circle"
        />
      </mask>
    </defs>
    <foreignObject x="0" y="0" width={diameter + strokeWidth} height={diameter + strokeWidth} mask={`url(#${id}-mask)`}>
      <div class="bg" />
    </foreignObject>
    <circle
      style={`animation-play-state:${isRunning ? 'running' : 'paused'}`}
      cx={diameter + strokeWidth / 2}
      cy={(diameter + strokeWidth) / 2}
      r={strokeWidth / 2 - 0.2}
      class="circle leading"
    />
  </svg>
</div>

<style>
  #container {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    --animation-shorthand: var(--duration) var(--animation-easing) 0s var(--animation-iteration-count)
      var(--animation-direction);
  }
  svg {
    fill: none;
    transform: rotate(-90deg);
  }

  .bg {
    background: conic-gradient(var(--color), hsl(from var(--color) calc(h - 15) s l));
    width: 100%;
    height: 100%;
    rotate: 90deg;
  }

  .background-ring,
  .progress-ring {
    stroke-width: var(--stroke-width);
  }

  .background-ring {
    stroke: rgba(255, 255, 255, 0.2);
  }

  .progress-ring {
    stroke: white;
    stroke-dasharray: var(--max-dash-length);
    stroke-dashoffset: var(--max-dash-length);
    animation: var(--animation-shorthand) spin;
  }
  .circle {
    fill: rgb(255, 255, 255);
    transform-origin: 50% 50%;
    animation: var(--animation-shorthand) rotateClockwise;
  }

  .leading {
    fill: hsl(from var(--color) calc(h - 15) s l);
    filter: drop-shadow(0px 4px 1px rgb(0 0 0 / 0.3));
    animation:
      var(--animation-shorthand) rotateClockwise,
      var(--animation-shorthand) colorChange;
  }

  @keyframes spin {
    0% {
      stroke-dashoffset: var(--max-dash-length);
    }
    100% {
      stroke-dashoffset: 0;
    }
  }

  @keyframes rotateClockwise {
    100% {
      transform: rotate(360deg);
    }
  }

  @keyframes colorChange {
    0% {
      opacity: 1;
    }
    2% {
      opacity: 0;
    }
    82% {
      opacity: 0;
    }
    83% {
      opacity: 1;
    }
  }
</style>
