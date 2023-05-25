(function () {
    function postToServerPerform(data) {

        // var endpoint;
        // if (type === 'static') {
        //     endpoint = '/json/static';
        // } else if (type === 'performance') {
        //     endpoint = '/json/performance';
        // } else if (type === 'activity') {
        //     endpoint = '/json/activity';
        // }

        fetch('json/performance', { // replace with your server endpoint
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        }).catch(error => console.error('Error:', error));
    }

    function postToServerStatic(data) {

        // var endpoint;
        // if (type === 'static') {
        //     endpoint = '/json/static';
        // } else if (type === 'performance') {
        //     endpoint = '/json/performance';
        // } else if (type === 'activity') {
        //     endpoint = '/json/activity';
        // }

        fetch('json/static', { // replace with your server endpoint
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        }).catch(error => console.error('Error:', error));
    }

    function postToServerActive(data) {

        // var endpoint;
        // if (type === 'static') {
        //     endpoint = '/json/static';
        // } else if (type === 'performance') {
        //     endpoint = '/json/performance';
        // } else if (type === 'activity') {
        //     endpoint = '/json/activity';
        // }

        fetch('json/activity', { // replace with your server endpoint
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(data),
        }).catch(error => console.error('Error:', error));
    }

    var Collector = {
        data: {
            staticData: {},
            performanceData: {},
            activityData: {
                errors: [],
                mouseActivity: [],
                keyboardActivity: [],
                idleTimes: [],
                session: {},
            },
        },
        start: function() {
            this.collectStaticData();
            this.collectPerformanceData();
            this.collectActivityData();
        },
        collectStaticData: function() {
            this.data.staticData.userAgent = navigator.userAgent;
            this.data.staticData.language = navigator.language;
            this.data.staticData.cookiesEnabled = navigator.cookieEnabled;
            this.data.staticData.javaScriptEnabled = true;
            this.data.staticData.screenDimensions = window.screen.width + 'x' + window.screen.height;
            this.data.staticData.windowDimensions = window.innerWidth + 'x' + window.innerHeight;
            this.data.staticData.connection = navigator.connection ? navigator.connection.effectiveType : 'unknown';
        },
        collectPerformanceData: function() {
            var timing = performance.timing;
            this.data.performanceData.timing = timing;
            this.data.performanceData.pageLoadStart = timing.navigationStart;
            this.data.performanceData.pageLoadEnd = timing.loadEventEnd;
            this.data.performanceData.totalLoadTime = timing.loadEventEnd - timing.navigationStart;
        },
        collectActivityData: function() {


            var idleStart = null;
            var idleEnd = null;
            var idleTimeout = null;
            const idleDelay = 2000; // idle time threshold in milliseconds

            function resetIdleTimer() {
                if (idleStart !== null) {
                    idleEnd = Date.now();
                    Collector.data.activityData.idleTimes.push({
                        start: idleStart,
                        end: idleEnd,
                        duration: idleEnd - idleStart
                    });
                    idleStart = null;
                    idleEnd = null;
                }
                clearTimeout(idleTimeout);
                idleTimeout = setTimeout(() => {
                    idleStart = Date.now();
                }, idleDelay);
            }



            window.onerror = function (message, url, line, column, error) {
                Collector.data.activityData.errors.push({
                    message: message,
                    url: url,
                    line: line,
                    column: column,
                    error: error
                });
            };

            window.addEventListener('mousemove', function(e) {
                resetIdleTimer();
                Collector.data.activityData.mouseActivity.push({
                    type: 'mousemove',
                    coordinates: { x: e.clientX, y: e.clientY },
                    timestamp: Date.now()
                });
            });

            window.addEventListener('click', function(e) {
                resetIdleTimer();
                Collector.data.activityData.mouseActivity.push({
                    type: 'click',
                    button: e.button,
                    coordinates: { x: e.clientX, y: e.clientY },
                    timestamp: Date.now()
                });
            });

            window.addEventListener('scroll', function() {
                resetIdleTimer();
                Collector.data.activityData.mouseActivity.push({
                    type: 'scroll',
                    coordinates: { x: window.scrollX, y: window.scrollY },
                    timestamp: Date.now()
                });
            });

            window.addEventListener('keydown', function(e) {
                resetIdleTimer();
                Collector.data.activityData.keyboardActivity.push({
                    type: 'keydown',
                    key: e.key,
                    timestamp: Date.now()
                });
            });

            window.addEventListener('keyup', function(e) {
                resetIdleTimer();
                Collector.data.activityData.keyboardActivity.push({
                    type: 'keyup',
                    key: e.key,
                    timestamp: Date.now()
                });
            });
            

            // var idleStart = null;
            // var idleEnd = null;

            // var activityObserver = new IdleDetector();
            // activityObserver.addEventListener('change', () => {
            //     if (activityObserver.userState === 'idle') {
            //         idleStart = Date.now();
            //     } else {
            //         idleEnd = Date.now();
            //         Collector.data.activityData.idleTimes.push({
            //             start: idleStart,
            //             end: idleEnd,
            //             duration: idleEnd - idleStart
            //         });
            //         idleStart = null;
            //         idleEnd = null;
            //     }
            // });
            // activityObserver.start();

            this.data.activityData.session.start = Date.now();
            this.data.activityData.session.page = window.location.href;

            window.addEventListener('beforeunload', function() {
                Collector.data.activityData.session.end = Date.now();
                Collector.data.activityData.session.duration = Collector.data.activityData.session.end - Collector.data.activityData.session.start;
                // Send data to your server
                postToServerPerform(Collector.data.performanceData);
                postToServerStatic(Collector.data.staticData)
                postToServerActive(Collector.data.activityData);
                // postToServer(Collector.data.staticData, 'static');
                // postToServer(Collector.data.performanceData, 'performance');
                // postToServer(Collector.data.performanceData, 'activity');
            });
        }
    };

    Collector.start();
})();