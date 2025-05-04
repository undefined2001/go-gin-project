let score = 0;
let totalQuestions = document.querySelectorAll('.question-block').length;

function checkAnswer(input) {
    const block = input.closest('.question-block');

    if (block.classList.contains("answered")) return; // prevent double-answer

    const correctAnswer = block.dataset.correctAnswer;
    const selectedAnswer = input.value;
    const allRadios = block.querySelectorAll('input[type="radio"]');

    // Lock inputs
    allRadios.forEach(r => r.disabled = true);
    block.classList.add("answered");

    if (selectedAnswer === correctAnswer) {
        score++;
        input.parentElement.classList.add('bg-success', 'text-white', 'rounded');
    } else {
        input.parentElement.classList.add('bg-danger', 'text-white', 'rounded');

        const correctRadio = [...allRadios].find(r => r.value === correctAnswer);
        if (correctRadio) {
            correctRadio.parentElement.classList.add('bg-success', 'text-white', 'rounded');
        }
    }
}

function showFinalScore() {
    const answered = document.querySelectorAll('.question-block.answered').length;
    if (answered < totalQuestions) {
        alert(`Please answer all questions before viewing your score.`);
        return;
    }
    alert(`Your Score: ${score} / ${totalQuestions}`);
}