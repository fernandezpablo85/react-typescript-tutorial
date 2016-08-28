var __extends = (this && this.__extends) || function (d, b) {
    for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p];
    function __() { this.constructor = d; }
    d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
};
/// <reference path='all.d.ts' />
var CommentBox = (function (_super) {
    __extends(CommentBox, _super);
    function CommentBox() {
        _super.call(this);
        this.state = { data: [] };
    }
    CommentBox.prototype.fetchComments = function () {
        var _this = this;
        $.ajax({
            url: this.props.url,
            dataType: 'json',
            cache: false,
            success: function (data) { return _this.setState({ data: data }); },
            error: function (xhr, status, err) { return console.error(status, err); }
        });
    };
    CommentBox.prototype.componentDidMount = function () {
        this.fetchComments();
        setInterval(this.fetchComments.bind(this), this.props.pollInterval);
    };
    CommentBox.prototype.render = function () {
        var _this = this;
        var handleCommentSubmit = function (comment) {
            console.warn('comment submitted!', comment);
            var updated = _this.state.data.slice(0);
            updated.push(comment);
            _this.setState({ data: updated });
        };
        return (React.createElement("div", {className: "commentBox"}, React.createElement("h1", null, "Comments"), React.createElement(CommentList, {data: this.state.data}), React.createElement(CommentForm, {onCommentSubmit: handleCommentSubmit})));
    };
    return CommentBox;
}(React.Component));
var CommentList = (function (_super) {
    __extends(CommentList, _super);
    function CommentList() {
        _super.apply(this, arguments);
    }
    CommentList.prototype.render = function () {
        var nodes = this.props.data.map(function (comment) { return React.createElement(Commentt, {author: comment.author}, comment.text); });
        return (React.createElement("div", {className: "commentList"}, nodes));
    };
    return CommentList;
}(React.Component));
;
var CommentForm = (function (_super) {
    __extends(CommentForm, _super);
    function CommentForm() {
        _super.apply(this, arguments);
    }
    CommentForm.prototype.render = function () {
        var _this = this;
        var handleSubmit = function (e) {
            e.preventDefault();
            var author = _this.refs['author'].value.trim();
            var text = _this.refs['text'].value.trim();
            if (author.length > 0 && text.length > 0) {
                _this.props.onCommentSubmit({ author: author, text: text });
            }
        };
        return (React.createElement("form", {className: "commentForm", onSubmit: handleSubmit}, React.createElement("input", {type: "text", placeholder: "Your name", ref: "author"}), React.createElement("input", {type: "text", placeholder: "Say something...", ref: "text"}), React.createElement("input", {type: "submit", value: "Post"})));
    };
    return CommentForm;
}(React.Component));
;
var Commentt = (function (_super) {
    __extends(Commentt, _super);
    function Commentt() {
        _super.apply(this, arguments);
    }
    Commentt.prototype.render = function () {
        return (React.createElement("div", {className: "comment"}, React.createElement("h2", {className: "commentAuthor"}, this.props.author), this.props.children));
    };
    return Commentt;
}(React.Component));
ReactDOM.render(React.createElement(CommentBox, {url: "/public/js/comments.json", pollInterval: 5000}), document.getElementById('content'));
