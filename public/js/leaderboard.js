/**
 * Leaderboard interactions (e.g. challenging, accept/decline),
 * interface data updates and navigation control
 */
(function(window, document, undefined) {

  /**
   * @constuctor
   */
	var Leaderboard = function(userId) {
    this.userId = userId;
    this.rootDomain = this.getRootDomain();

    this.restorePageFromUrl();
    this.bindGotoPage();
    this.bindChallengeUser();
    this.bindUserChallenges();
    this.getUserChallenges();
  };
  window.Leaderboard = Leaderboard;

  /**
   * Handle accepting/declining challenges
   */
  Leaderboard.prototype.bindUserChallenges = function(){
        // bind accept/decline challgenge
  };

  /**
   * Handle fetching the user challenges every x seconds
   */
  Leaderboard.prototype.getUserChallenges = function(){
    var userId = this.userId;

    var poll = function() {
      $.get('challenges/' + userId, function(data) {
        $('#user-challenges').html(data);
        setTimeout(poll, 10000);
      });
    };
    poll();
  };

  /**
   * Fetch the list of leaderboards
   */
  Leaderboard.prototype.getAllLeaderboards = function() {
    $.get('leaderboards/view', function(data) {

    });
  };

  /**
   * Get the root domain of a URL
   */
  Leaderboard.prototype.getRootDomain = function() {
    var href = window.location.href,
        domain = href.split('//').splice(1)[0].split('/')[0]

    return domain;
  };

  /** 
   * Restore a pages state based on the URL
   */
  Leaderboard.prototype.restorePageFromUrl = function() {
    var self = this,
        href = window.location.href,
        currentPage = href.split('#').splice(1).join();

    if (currentPage.length > 0) {
      this.gotoPage(currentPage);
    } else {
      this.gotoPage('leaderboards/view');
    }

    // Handle back and forward refresh
    $(window).on('popstate', function() {
      self.restorePageFromUrl();
    });
  };

  /**
   * Prevent default behavior of links,
   * fetch the href via AJAX and dump into #page-content
   */
  Leaderboard.prototype.bindGotoPage = function() {
    var self = this;
    $(document).on('click', 'a:not(.no-ajax)', function(e) {
      e.preventDefault();

      var $self = $(this);
      var href = $(this).attr("href");

      self.gotoPage(href);
    });
  };

  /**
   * Send a user to a particular page via AJAX, update the content
   * and update the browser history
   */
  Leaderboard.prototype.gotoPage = function(href) {
    var self = this;

    $.get(href, function(data) {
      $('#page-content').html(data);
      window.history.pushState({ url: href }, '', '#' + href);
    })
    .fail(function() {
      self.gotoPage('leaderboards/view');
    });
  };

  Leaderboard.prototype.bindChallengeUser = function() {
    $(document).on('click', '.js-challenge', function() {
      var $self = $(this);

      var boardId = $self.data('board-id'),
          toUserId = $self.data('user-id'),
          fromUserId = $('#from-user-id').val();

      $.post('/challenge/create', 
        { 
          leaderBoardId : boardId, 
          toUserId : toUserId, 
          fromUserId : fromUserId
        }, 
        function(data) {
          alert(data);
      });
    });
  }

})(window, document)