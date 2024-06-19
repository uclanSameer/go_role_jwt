provider "aws" {
    region = "us-west-2"
}

resource "aws_instance" "go_role_based_auth" {
    ami           = "ami-0c94855ba95c574c8" # replace with your AMI ID
    instance_type = "t2.micro"

    tags = {
        Name = "go_role_based_auth"
    }

    provisioner "file" {
        source      = "path/to/your/go/executable"
        destination = "/path/on/the/instance"
    }

    provisioner "remote-exec" {
        inline = [
            "chmod +x /path/on/the/instance",
            "/path/on/the/instance"
        ]
    }

    connection {
        type        = "ssh"
        user        = "ec2-user"
        private_key = file("~/.ssh/id_rsa")
        host        = self.public_ip
    }
}

