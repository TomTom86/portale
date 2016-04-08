




    <div class="modal-dialog">
                    <div class="loginmodal-container">
                    <div id="brand">&nbsp</div>
                    &nbsp;
                    {{if .flash.error}}
                    <div class="alert alert-danger">
                        <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                            ×</button>
                        <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                        <hr class="message-inner-separator">
                        {{.flash.error}}.
                    </div>
                    {{end}}
                    {{if .Errors}}
                    {{range $rec := .Errors}}
                    <div class="alert alert-danger">
                        <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                            ×</button>
                        <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                        <hr class="message-inner-separator">
                        {{$rec}}.
                    </div>      
                    {{end}}
                    &nbsp;
                    {{end}}
                    <form method="POST">
                        {{.xsrfdata}}
                        <input type="text" name="email" placeholder="Username" required="">
                        <input type="password" name="password" placeholder="Password" required="">
                        <input type="submit" name="login" value="Login" class="login loginmodal-submit" value="Login">
                    </form>
                        
                    <div class="login-help">
                        <a href="http://{{.domainname}}/register">Registrati</a> - <a href="http://{{.domainname}}/forgot">Password dimenticata</a>
                    </div>
                    </div>
                </div>
            </div>
    </div>




