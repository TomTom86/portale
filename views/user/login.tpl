




    <div class="modal-dialog">
                    <div class="loginmodal-container">
                    <div id="brand">&nbsp</div>
                    &nbsp;
                    {{if .flash.error}}
                    <h3>{{.flash.error}}</h3>
                    &nbsp;
                    {{end}}
                    {{if .Errors}}
                    {{range $rec := .Errors}}
                    <h3>{{$rec}}</h3>
                    {{end}}
                    &nbsp;
                    {{end}}
                        &nbsp
                    <form method="POST">
                        {{.xsrfdata}}
                        <input type="text" name="email" placeholder="Username">
                        <input type="password" name="password" placeholder="Password">
                        <input type="submit" name="login" value="Login" class="login loginmodal-submit" value="Login">
                    </form>
                        
                    <div class="login-help">
                        <a href="http://{{.domainname}}/register">Registrati</a> - <a href="http://{{.domainname}}/forgot">Password dimenticata</a>
                    </div>
                    </div>
                </div>
            </div>
    </div>




