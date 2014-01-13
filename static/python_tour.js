$(function() {
    icon = $('#lang-icon')
    icon.attr("data-content", "正确")
    icon.popover("show")
    window.setTimeout(function() {
        icon.popover("hide")
    }, 3000);
});
