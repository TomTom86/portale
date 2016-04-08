<div class="col-md-10 content">
    <div class="panel panel-default">
        <div class="panel-heading">
            <h1>{{.First}} {{.Last}}</h1>
        </div>
        <div class="panel-body">
            {{if .flash.error}}
            <div class="alert alert-danger">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
                <span class="glyphicon glyphicon-hand-right"></span> <strong> ATTENZIONE - Errore</strong>
                <hr class="message-inner-separator">
                {{.flash.error}}.
            </div>
            {{end}}
            {{if .flash.notice}}
            <div class="alert alert-success">
                <button type="button" class="close" data-dismiss="alert" aria-hidden="true">
                    ×</button>
               <span class="glyphicon glyphicon-ok"></span> <strong>Operazione completata con successo</strong>
                <hr class="message-inner-separator">
                {{.flash.notice}}.
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
            <form class="form-horizontal" method="POST">
                {{.xsrfdata}}
                <!-- Nome input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="name">Nome</label>  
                <div class="col-md-4">
                <input id="first" name="first" type="text" value="{{.First}}" class="form-control input-md">
                    
                </div>
                </div>

                <!-- Cognome input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="last">Cognome</label>  
                <div class="col-md-4">
                <input id="last" name="last" type="text" value="{{.Last}}" class="form-control input-md">
                    
                </div>
                </div>

                <!-- Email input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="email">Email</label>  
                <div class="col-md-4">
                <input id="email" name="email" type="text" value="{{.Email}}" class="form-control input-md">
                    
                </div>
                </div>

                <!-- Password input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="current">Password Attuale</label>
                <div class="col-md-4">
                    <input id="current" name="current" type="password" placeholder="Password" class="form-control input-md" required="">
                    <span class="help-block">Inserisci la tua password</span>
                </div>
                </div>
                    <hr>

                <!-- NewPassword input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="password">Nuova Password</label>
                <div class="col-md-4">
                    <input id="password" name="password" type="password" placeholder="Password" class="form-control input-md">
                    <span class="help-block">Inserisci una password di almeno 6 caratteri</span>
                </div>
                </div>

                <!-- Conferma NewPassword input-->
                <div class="form-group">
                <label class="col-md-4 control-label" for="password2">Ripeti Nuova Password</label>
                <div class="col-md-4">
                    <input id="password2" name="password2" type="password" placeholder="Password" class="form-control input-md">
                    <span class="help-block">Inserisci nuovamente la nuova password</span>
                </div>
                </div>

                <br/>
                <button class="btn btn-primary" onclick="location.href='http://localhost:8080/user/remove'">Cancella Account</button>
                <button class="btn btn-primary pull-right" type="submit" class="pull-right" value="Aggiorna" />Aggiorna</button>
                    
            </form>                                              
        </div>
    </div>

</div>
          