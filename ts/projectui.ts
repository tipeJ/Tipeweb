window.onload = () => {
    let newParent = document.createElement('div');
    newParent.id = 'post-content';
    
    Array.prototype.forEach.call(document.querySelectorAll('.outline-2'), function(c: Element){
        newParent.appendChild(c);
    });
    document.getElementById('content').appendChild(newParent);
};
