<nav class="navbar navbar-default navbar-fixed-top">
    <div class="brand" style="padding: 18px 36px;">
        <a href="index.html">
            <img src="/static/img/logo.png" width="139" height="40" alt="eWorkBook" class="img-responsive logo"></a>
        </div>
        <div class="container-fluid">
            <div class="navbar-btn">
                <button type="button" class="btn-toggle-fullwidth">
                    <i class="lnr lnr-arrow-left-circle"></i>
                </button>
            </div>
            <form class="navbar-form navbar-left">
                <div class="input-group">
                    <input type="text" value="" class="form-control" placeholder="Search dashboard...">
                        <span class="input-group-btn">
                            <button type="button" class="btn btn-primary">Go</button>
                        </span>
                    </div>
                </form>

                <div id="navbar-menu">
                    <ul class="nav navbar-nav navbar-right">
                        <li class="dropdown">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <i class="lnr lnr-question-circle"></i>
                                <span>Help</span>
                                <i class="icon-submenu lnr lnr-chevron-down"></i>
                            </a>
                            <ul class="dropdown-menu">
                                <li>
                                    <a href="#">Basic Use</a>
                                </li>
                                <li>
                                    <a href="#">Working With Data</a>
                                </li>
                                <li>
                                    <a href="#">Security</a>
                                </li>
                                <li>
                                    <a href="#">Troubleshooting</a>
                                </li>
                            </ul>
                        </li>
                        <li class="dropdown">
                            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
                                <img src="/static/img/user.png" class="img-circle" alt="Avatar">
                                    <span>Samuel</span>
                                    <i class="icon-submenu lnr lnr-chevron-down"></i>
                                </a>
                                <ul class="dropdown-menu">
                                    <li>
                                        <a href="profile">
                                            <i class="lnr lnr-user"></i>
                                            <span>My Profile</span>
                                        </a>
                                    </li>
                                    <li>
                                        <a href="#">
                                            <i class="lnr lnr-envelope"></i>
                                            <span>Message</span>
                                        </a>
                                    </li>
                                    <li>
                                        <a href="#">
                                            <i class="lnr lnr-cog"></i>
                                            <span>Settings</span>
                                        </a>
                                    </li>
                                    <li>
                                        <a href="logout">
                                            <i class="lnr lnr-exit"></i>
                                            <span>Logout</span>
                                        </a>
                                    </li>
                                </ul>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
            <!-- END NAVBAR -->
            <!-- LEFT SIDEBAR -->
            <div id="sidebar-nav" class="sidebar">
                <div class="sidebar-scroll">
                    <nav>
                        <ul class="nav">
                            <li>
                                <a href="{{ urlfor "MainController.Get"}}" {{if .index}} class="active" {{end}}>
                                    <i class="lnr lnr-home"></i>
                                    <span>Home</span>
                                </a>
                            </li>
                            <li>
                                <a href="{{ urlfor "SearchController.GetSearch"}}" {{if .searchMenu}} class="active" {{end}}>
                                    <i class="lnr lnr-code"></i>
                                    <span>Search</span>
                                </a>
                            </li>
                            <li>
                                <a href="{{ urlfor "UserController.ListUsers"}}" {{if .usersMenu}} class="active" {{end}}>
                                    <i class="lnr lnr-user"></i>
                                    <span>Users</span>
                                </a>
                            </li>
                            <li>
                                <a href="{{ urlfor "TeamController.ListTeams"}}" {{if .teamsMenu}} class="active" {{end}}>
                                    <i class="lnr lnr-users"></i>
                                    <span>Teams</span>
                                </a>
                            </li>
                            <li>
                                <a href="{{ urlfor "PodController.ListPods"}}" {{if .podsMenu}} class="active" {{end}}>
                                    <i class="lnr lnr-shirt"></i>
                                    <span>POD</span>
                                </a>
                            </li>
                            <li>
                                <a href="{{ urlfor "ProjectController.ListProjects"}}" {{if .projectMenu}} class="active" {{end}}>
                                    <i class="lnr lnr-diamond"></i>
                                    <span>Project</span>
                                </a>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
            <!-- END LEFT SIDEBAR -->
