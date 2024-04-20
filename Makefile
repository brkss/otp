

NAME = ft_otp



all:
	go build -o $(NAME)


$(NAME): all

clean:
	rm -Rf $(NAME)

fclean: clean
