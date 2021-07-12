$(function(){

    //首页nav
    $(".nav_list li").hover(function () {
        $(this).find('ul').stop(false,true).slideDown();
    },function () {
        $(this).find('ul').stop(false,true).slideUp();
    });


    /*手机左侧nav*/
    $("ul.leftOne a").click(function (event) {
        if ($(this).siblings("ul").length > 0) {
            var a = $(this);
            var thisname = a.attr('class');
            if (thisname == null || thisname == 0) {
                a.siblings("ul").slideDown(300);
                a.parent().siblings().find('ul').slideUp(500).siblings("a").removeClass();
                var parent = a.parent().parents("ul").attr('class');
                switch (parent) {
                    case "leftOne":
                        a.addClass('box-1');
                        break;
                    case "two":
                        a.addClass('box-2');
                        a.find('i').text('-').parent().parent().siblings().find('i').text('+');
                        break;
                    case "three":
                        a.parent().addClass('box-3').siblings().removeClass('box-3');
                        break;
                }
            } else {
                a.removeClass().siblings('ul').slideUp(300);
                a.siblings('ul').find('a').removeClass().siblings('ul').slideUp(300);
                a.find('i').text('+');
            };
            return false;
        };
    });
    $('.menu').click('touchstart',function () {
        $('.leftNav').addClass('on');
    });

    $(".leftNavClose,.leftNavBg").click('touchstart',function () {
        $(this).parents(".leftNav").removeClass("on");
    });

    //banner大图轮播
    $(".banSlider").slick({
        dots:true//指示点
    });

    $(".newsSlider").slick({
        infinite:false,//循环
        dots:true,
        slidesToShow:3,//显示个数
        slidesToScroll:3,//轮播个数
        responsive:[
            {
                breakpoint:767,
                settings:{
                    slidesToShow: 1,
                    slidesToScroll: 1
                }
            }
        ]
    });
    $(".linkSlider").slick({
        infinite:false,//循环
        dots:false,
        arrows:true,
        // autoplay:true,
        // autoplaySpeed:2000,//自动播放间隔
        slidesToShow:8,//显示个数
        slidesToScroll:2,//轮播个数
        responsive:[
            {
                breakpoint:992,
                settings:{
                    slidesToShow: 6,
                    slidesToScroll: 2
                }
            },
            {
                breakpoint:767,
                settings:{
                    slidesToShow: 3,
                    slidesToScroll: 1
                }
            }
        ]
    });
    // 滚回顶部
    $(".scroll_top").click(function(){
        $("html,body").animate({scrollTop:'0px'},800);
    });
    $(window).scroll(function () {
        var top = $(".scroll_top");
        if ($(window).scrollTop() == $(document).height() - $(window).height()) {
            top.addClass("on");
        }else{
            top.removeClass("on");
        }
    });


    //内页左侧分类
    $("ul.detNav a").click(function (event) {
        if ($(this).siblings("ul").length > 0) {
            var a = $(this);
            var thisname = a.attr('class');
            if (thisname == null || thisname == 0) {
                a.siblings("ul").slideDown(300);
                a.parent().siblings().find('ul').slideUp(500).siblings("a").removeClass();
                var parent = a.parent().parents("ul").attr('class');
                switch (parent) {
                    case "detNav":
                        a.addClass('box-on');
                        a.find('i').addClass('in').parent().parent().siblings().find('i').removeClass('in');
                        a.find('i').text('-').parent().parent().siblings().find('i').text('+');
                        break;
                    case "detNav-two":
                        a.parent().addClass('box-on').siblings().removeClass('box-on');
                        break;
                    case "detNav-three":
                        a.addClass('box-on');
                        break;
                }
            } else {
                a.removeClass().siblings('ul').slideUp(300);
                a.siblings('ul').find('a').removeClass().siblings('ul').slideUp(300);
                a.find("i").removeClass('in');
                a.find('i').text('+');
            };
            return false;
        };
    });

    //手机内页分类
    $(".det-leftTitle em").click(function () {
        $(".detNav").slideToggle();
        $(this).toggleClass('on');
        if($(this).hasClass('on')){
            $(this).text('-');
        }else{
            $(this).text('+')
        }
    });



});



