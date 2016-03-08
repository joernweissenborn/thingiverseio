#include<stdio.h>
#include "thingiverseio.h"
#include <windows.h>

  char * const DESCRIPTOR = "\
  functions:\n\
    - name: SayHello\n\
      input:\n\
        - name: Greeting\n\
          type: string\n\
      output:\n\
        - name: Answer\n\
          type: string\n\
  ";

  int main() {

    printf("Testing Input Creation...\n");

    int input = tvio_new_input(DESCRIPTOR);

	if (input == -1) {
		printf("FAIL\n");
		return 1;
	};
	printf("SUCCES\n");

	printf("Testing Output Creation...\n");

	int output = tvio_new_output(DESCRIPTOR);

	if (output == -1) {
		printf("FAIL\n");
		return 1;
	};
	printf("SUCCES\n");

	printf("Testing Call...\n");

	char * uuid;
	int uuid_size;

	char * fun = "SayHello";
	char * params = "HELLO";
	int params_size = 5;

	int err = tvio_call(input, fun,params,params_size, &uuid, &uuid_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (uuid_size == 0) {
		printf("FAIL, uuid_size is 0\n");
		return 1;
	};

	sleep(5);

	char * req_uuid;
	int req_uuid_size;
	err = tvio_get_next_request_id(output, &req_uuid, &req_uuid_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (req_uuid_size == 0) {
		printf("FAIL, req_uuid_size is 0\n");
		return 1;
	};

	char * rfun;
	int rfun_size;
	err = tvio_retrieve_request_function(output, uuid, &rfun, &rfun_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (rfun_size == 0) {
		printf("FAIL, fun_size is 0\n");
		return 1;
	};
	char * rparams;
	int rparams_size;
	err = tvio_retrieve_request_params(output, uuid, &rparams, &rparams_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (rparams_size != 5) {
		printf("FAIL, rparams_size is 0\n");
		return 1;
	};

	char * resparams = "HELLO_BACK";
	int resparams_size = 10;

	err = tvio_reply(output, uuid, resparams, resparams_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};

	Sleep(5);

	int ready;
	err = tvio_result_ready(input, uuid, &ready);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (ready != 1) {
		printf("FAIL, result hasnt arrived\n");
		return 1;
	}

	char * resultparams;
	int resultparams_size;
	err = tvio_retrieve_result_params(input, uuid, &resultparams, &resultparams_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (resultparams_size != 10) {
		printf("FAIL, rparams_size is 0\n");
		return 1;
	};

	printf("SUCCES\n");

	printf("Testing Trigger...\n");

	err = tvio_start_listen(input, fun);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};

	Sleep(5);

	err = tvio_trigger(input, fun,params,params_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	Sleep(5);

	err = tvio_get_next_request_id(output, &req_uuid, &req_uuid_size);
	if (err != 0) {
		printf("FAIL, get_gext_req err not 0\n");
		return 1;
	};
	if (req_uuid_size == 0) {
		printf("FAIL, req_uuid_size is 0\n");
		return 1;
	};

	err = tvio_reply(output, req_uuid, resparams, resparams_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};

	Sleep(5);

	err = tvio_listen_result_available(input, &ready);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (ready != 1) {
		printf("FAIL, result hasnt arrived\n");
		return 1;
	}

	err = tvio_retrieve_listen_result_params(input, &resultparams, &resultparams_size);
	if (err != 0) {
		printf("FAIL, err not 0\n");
		return 1;
	};
	if (resultparams_size != 10) {
		printf("FAIL, rparams_size is 0\n");
		return 1;
	};

	printf("SUCCES\n");

	return 0;
}
