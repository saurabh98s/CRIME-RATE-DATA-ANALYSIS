var ctx = document.getElementById('myChart');
var chart = new Chart(ctx, {
    // The type of chart we want to create
    type: 'line',

    // The data for our dataset
    data: {
        labels: ['Jharkhand', 'Bihar', 'UP', 'Rajsthan'],
        datasets: [{
            label: 'Auto theft stolen',
            backgroundColor: 'rgb(255, 99, 132)',
            borderColor: 'rgb(255, 99, 132)',
            data: [1120, 1110, 895, 2212, 3220]
        }]
    },

    // Configuration options go here
    options: {}
});