window.onload = function () {
    var newParent = document.createElement('div');
    newParent.id = 'post-content';
    Array.prototype.forEach.call(document.querySelectorAll('.outline-2'), function (c) {
        newParent.appendChild(c);
    });
    document.getElementById('content').appendChild(newParent);
};
//# sourceMappingURL=projectui.js.map