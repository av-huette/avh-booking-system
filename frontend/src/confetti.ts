import confetti from 'canvas-confetti'
import {ref} from 'vue'

// Simple confetti function
export const fireConfetti = () => {
    confetti({
        particleCount: 100,
        spread: 70,
        origin: {y: 0.6}
    })
} // fireConfetti

// Confetti from both sides
export const fireConfettiFromSides = () => {
    const end = Date.now() + (3 * 1000)
    const colors = ['#9fc5e8', '#ffffff']

    ;(function frame() {
        confetti({
            particleCount: 2,
            angle: 60,
            spread: 55,
            origin: {x: 0},
            colors: colors
        })
        confetti({
            particleCount: 2,
            angle: 120,
            spread: 55,
            origin: {x: 1},
            colors: colors
        })

        if (Date.now() < end) {
            requestAnimationFrame(frame)
        }
    }())
} // fireConfettiFromSides

// Fireworks
export const fireworks = () => {
    var duration = 5 * 1000;
    var animationEnd = Date.now() + duration;
    var defaults = {startVelocity: 30, spread: 360, ticks: 60, zIndex: 0};

    function randomInRange(min: number, max: number) {
        return Math.random() * (max - min) + min;
    }

    var interval = setInterval(function () {
        var timeLeft = animationEnd - Date.now();

        if (timeLeft <= 0) {
            return clearInterval(interval);
        }

        var particleCount = 50 * (timeLeft / duration);
        // since particles fall down, start a bit higher than random
        confetti({...defaults, particleCount, origin: {x: randomInRange(0.1, 0.3), y: Math.random() - 0.2}});
        confetti({...defaults, particleCount, origin: {x: randomInRange(0.7, 0.9), y: Math.random() - 0.2}});
    }, 250);
} // fireworks

// Snow effect
let snowAnimation: number | null = null
export const isSnowing = ref(false)

export const startSnow = () => {
    if (isSnowing.value) {
        stopSnow()
        return
    }

    isSnowing.value = true
    const duration = 3 * 1000
    const animationEnd = Date.now() + duration
    let skew = 1

    const randomInRange = (min: number, max: number) => {
        return Math.random() * (max - min) + min
    }

    const frame = () => {
        const timeLeft = animationEnd - Date.now()
        var ticks = Math.max(200, 500 * (timeLeft / duration));
        skew = Math.max(0.8, skew - 0.001)

        // Create multiple snowflakes with different colors
        for (let i = 0; i < 3; i++) {
            confetti({
                particleCount: 1,
                startVelocity: 0,
                ticks: ticks,
                gravity: randomInRange(0.4, 0.6),
                origin: {
                    x: Math.random(),
                    y: (Math.random() * skew) - 0.2
                },
                colors: [Math.random() > 0.5 ? '#ffffff' : '#9fc5e8'],
                shapes: ['circle'],
                scalar: randomInRange(0.3, 1),
                drift: randomInRange(-0.4, 0.4)
            })
        }

        if (timeLeft > 0) {
            snowAnimation = requestAnimationFrame(frame)
        } else {
            isSnowing.value = false
            snowAnimation = null
        }
    }

    frame()
} // startSnow

export const stopSnow = () => {
    if (snowAnimation) {
        cancelAnimationFrame(snowAnimation)
        snowAnimation = null
        isSnowing.value = false
    }
} // stopSnow

// Payment
export const paymentConfetti = () => {
    var scalar = 3;
    var coin = confetti.shapeFromText({text: '🪙', scalar});
    var dollarSign = confetti.shapeFromText({text: '💲', scalar});

    var defaults = {
        spread: 360,
        ticks: 100,
        gravity: 0,
        decay: 0.96,
        startVelocity: 12,
        scalar
    };

    function shoot() {
        confetti({
            ...defaults,
            particleCount: 16,
            shapes: [coin],
        });

        confetti({
            ...defaults,
            particleCount: 10,
            //flat: true,
            shapes: [dollarSign],
        });

    }

    setTimeout(shoot, 0);
    //setTimeout(shoot, 100);
    //setTimeout(shoot, 200);
} // paymentConfetti

// Booking
export const bookingConfetti = () => {
    var scalar = 5;
    //var beers = confetti.shapeFromText({text: '🍻', scalar});
    var beer = confetti.shapeFromText({text: '🍺', scalar});

    var defaults = {
        spread: 0, // horizontal movement
        ticks: 200,
        gravity: 0, // horizontal movement
        decay: 0.97,
        startVelocity: 5,
        scalar,
        particleCount: 1,
        flat: true,
        shapes: [beer],
    };

    function shoot1() {
        confetti({
            ...defaults,
            angle: 360,
            origin: {x: 0.4},
        });
    }

    function shoot2() {
        confetti({
            ...defaults,
            angle: 180,
            origin: {x: 0.6},
        });
    }

    setTimeout(shoot1, 0);
    setTimeout(shoot2, 0);
    //setTimeout(shoot, 100);
    //setTimeout(shoot, 200);
} // paymentConfetti

export const huetteConfetti = () => {
    var end = Date.now() + (5 * 1000);
    
    // go Buckeyes!
    var colors = ['#73b7f6', '#ffffff'];
    
    (function frame() {
      confetti({
        particleCount: 2,
        angle: 60,
        spread: 55,
        origin: { x: 0 },
        colors: colors
      });
      confetti({
        particleCount: 2,
        angle: 120,
        spread: 55,
        origin: { x: 1 },
        colors: colors
      });
    
      if (Date.now() < end) {
        requestAnimationFrame(frame);
      }
    }());
}