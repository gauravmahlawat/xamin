---
title: "Current Affairs Quiz"
date: 2019-11-19T17:21:20+01:00
page: true
---

# Current affairs quiz

{{< quizdown >}}

---
primary_color: '#fc4c02'
secondary_color: lightgray
text_color: black
shuffle_questions: true
---

## who is prime minister of india?

> modi

- [x] Narendra modi
- [ ] gandhi
- [ ] nehru
- [ ] kejriwal

## who is chief minister of Delhi?

> kejri

- [ ] Narendra modi
- [ ] gandhi
- [ ] nehru
- [x] kejriwal

## Which country hosted the 2021 BRICS summit?

- [ ] China
- [x] India
- [ ] Russia
- [ ] Brazil

## What is the name of India's first indigenous COVID-19 vaccine?

- [ ] Covishield
- [ ] Pfizer-BioNTech
- [x] Covaxin
- [ ] Moderna

## Which Indian state announced the implementation of the "One Nation One Ration Card" scheme, allowing portability of ration cards across the country?

- [ ] Maharashtra
- [ ] Karnataka
- [ ] Tamil Nadu
- [x] Haryana

## Who is the current President of India as of August 2023?

- [ ] Narendra Modi
- [ ] Sonia Gandhi
- [x] Ram Nath Kovind
- [ ] Rahul Gandhi

## Which city is known as the "Pink City" of India?

- [ ] Jaipur
- [ ] Mumbai
- [ ] Kolkata
- [x] Jodhpur


{{< /quizdown >}}

{{< rawhtml >}}
<div id="timer-container">
  <button id="start-button">Start Timer</button>
  <div id="timer">00:00</div>
</div>

<script>
  var timer;
  var startTime;

  document.getElementById("start-button").addEventListener("click", function() {
    if (!startTime) {
      startTime = new Date().getTime();
      timer = setInterval(updateTimer, 1000);
    }
  });

  function updateTimer() {
    var currentTime = new Date().getTime();
    var elapsedTime = currentTime - startTime;

    var minutes = Math.floor(elapsedTime / (1000 * 60));
    var seconds = Math.floor((elapsedTime % (1000 * 60)) / 1000);

    document.getElementById("timer").textContent = ("0" + minutes).slice(-2) + ":" + ("0" + seconds).slice(-2);
  }
</script>

     <script 
     src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/quizdown.js">
  </script>
  <script 
      src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/extensions/quizdownKatex.js">
  </script>
  <script 
      src="https://cdn.jsdelivr.net/npm/quizdown@latest/public/build/extensions/quizdownHighlight.js">
  </script>
  <script>quizdown.register(quizdownHighlight).register(quizdownKatex).init()</script> 
{{< /rawhtml >}}