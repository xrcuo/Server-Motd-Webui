
$.ajax({
    type: "get",
    url: "https://iuxcn.cn/Api/UserInfo",
    async: true,
    success: function (data) {
        window.info = data;
        var showFPS = (function () {
            var requestAnimationFrame =
                window.requestAnimationFrame ||
                window.webkitRequestAnimationFrame ||
                window.mozRequestAnimationFrame ||
                window.oRequestAnimationFrame ||
                window.msRequestAnimationFrame ||
                function (callback) {
                    window.setTimeout(callback, 1000 / 60);
                };
            var e, pe, pid, fps, last, offset, step, appendFps;

            fps = 0;
            last = Date.now();
            step = function () {
                offset = Date.now() - last;
                fps += 1;
                if (offset >= 1000) {
                    last += offset;
                    appendFps(fps);
                    fps = 0;
                }
                requestAnimationFrame(step);
            };
            appendFps = function (fps) {
                var settings = {
                    timeout: 5000,
                    logError: true
                }
                var p = new Ping(settings);
                p.ping('', function (perr, ping) {
                    if (perr) {
                        $('#fps').html('<span style="float:right">' + fps + 'FPS</span><br/><span style="float:right">' + window.info.data.os + '</span><br/><span style="float:right;margin-top:1px;">' + window.info.data.browser + '</span><br/><span style="float:right;margin-top:1px;">' + window.info.data.location + '</span><br/><span style="float:right;margin-top:1px;">©2024 椛椛</span><br/><span style="float:right;margin-top:1px;"><a href="//beian.miit.gov.cn/" target="_blank">闽ICP备2022001129号-1</a></span>');
                    }
                    if (ping) {
                        $('#fps').html('<span style="float:right">' + fps + 'FPS</span><br/><span style="float:right">' + window.info.data.os + '</span><br/><span style="float:right;margin-top:1px;">' + window.info.data.browser + '</span><br/><span style="float:right;margin-top:1px;">' + window.info.data.location + ' ' + ping + 'ms</span><br/><span style="float:right;margin-top:1px;">©2024 椛椛</span><br/><span style="float:right;margin-top:1px;"><a href="//beian.miit.gov.cn/" target="_blank">闽ICP备2022001129号-1</a></span>');
                    }
                    console.log(fps + 'FPS ' + ping + 'ms');
                });
            };
            step();
        })();
    }
});
$.ajax({
    type: "post",
    url: "https://iuxcn.cn/Api/MusicList",
    async: true,
    data: {
        format: 'json',
        url: 'https://y.qq.com/n/yqq/playlist/2882976222.html',
    },
    success: function (res) {
        window.ap = new APlayer({
            container: document.getElementById('player'),
            autoplay: false,
            fixed: true,
            preload: 'auto',
            mutex: true,
            lrcType: 3,
            audio: res.data,
        });
        window.ap.lrc.show();
        window.ap.list.show();
        //window.ap.play();
    }
});
