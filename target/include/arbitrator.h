#ifndef arbitrator_bindings
#define arbitrator_bindings

#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

#define SSTORE_SENTRY_GAS 2300

#define COLD_ACCOUNT_GAS 2600

#define COLD_SLOAD_GAS 2100

#define LOG_TOPIC_GAS 375

#define LOG_DATA_GAS 8

#define COPY_WORD_GAS 3

#define KECCAK_256_GAS 30

#define KECCAK_WORD_GAS 6

#define GAS_QUICK_STEP 2

#define ADDRESS_GAS GAS_QUICK_STEP

#define BASEFEE_GAS GAS_QUICK_STEP

#define CHAINID_GAS GAS_QUICK_STEP

#define COINBASE_GAS GAS_QUICK_STEP

#define GASLIMIT_GAS GAS_QUICK_STEP

#define NUMBER_GAS GAS_QUICK_STEP

#define TIMESTAMP_GAS GAS_QUICK_STEP

#define GASLEFT_GAS GAS_QUICK_STEP

#define CALLER_GAS GAS_QUICK_STEP

#define CALLVALUE_GAS GAS_QUICK_STEP

#define GASPRICE_GAS GAS_QUICK_STEP

#define ORIGIN_GAS GAS_QUICK_STEP

#define OperatorCode_OPERATOR_COUNT 529

/**
 * For hostios that may return something.
 */
#define HOSTIO_INK 8400

/**
 * For hostios that include pointers.
 */
#define PTR_INK (13440 - HOSTIO_INK)

/**
 * For hostios that involve an API cost.
 */
#define EVM_API_INK 59673

/**
 * WASM page size, or 2^16 bytes.
 */
#define PAGE_SIZE (1 << 16)

#define ARBITRATOR_MACHINE_STATUS_RUNNING 0

#define ARBITRATOR_MACHINE_STATUS_FINISHED 1

#define ARBITRATOR_MACHINE_STATUS_ERRORED 2

#define ARBITRATOR_MACHINE_STATUS_TOO_FAR 3

#define GLOBAL_STATE_BYTES32_NUM 2

#define GLOBAL_STATE_U64_NUM 2

#define Machine_MAX_STEPS (1 << 43)

enum EvmApiStatus {
  EvmApiStatus_Success,
  EvmApiStatus_Failure,
};
typedef uint8_t EvmApiStatus;

enum UserOutcomeKind {
  UserOutcomeKind_Success,
  UserOutcomeKind_Revert,
  UserOutcomeKind_Failure,
  UserOutcomeKind_OutOfInk,
  UserOutcomeKind_OutOfStack,
};
typedef uint8_t UserOutcomeKind;

typedef struct Machine Machine;

typedef struct GoSliceData {
  const uint8_t *ptr;
  uintptr_t len;
} GoSliceData;

typedef struct RustBytes {
  uint8_t *ptr;
  uintptr_t len;
  uintptr_t cap;
} RustBytes;

typedef struct Bytes32 {
  uint8_t bytes[32];
} Bytes32;

/**
 * Information about an activated program.
 */
typedef struct StylusData {
  /**
   * Global index for the amount of ink left.
   */
  uint32_t ink_left;
  /**
   * Global index for whether the program is out of ink.
   */
  uint32_t ink_status;
  /**
   * Global index for the amount of stack space remaining.
   */
  uint32_t depth_left;
  /**
   * Gas needed to invoke the program.
   */
  uint32_t init_gas;
  /**
   * Canonical estimate of the asm length in bytes.
   */
  uint32_t asm_estimate;
  /**
   * Initial memory size in pages.
   */
  uint16_t footprint;
  /**
   * Entrypoint offset.
   */
  uint32_t user_main;
} StylusData;

typedef struct PricingParams {
  /**
   * The price of ink, measured in bips of an evm gas
   */
  uint32_t ink_price;
} PricingParams;

typedef struct StylusConfig {
  /**
   * Version the program was compiled against
   */
  uint16_t version;
  /**
   * The maximum size of the stack, measured in words
   */
  uint32_t max_depth;
  /**
   * Pricing parameters supplied at runtime
   */
  struct PricingParams pricing;
} StylusConfig;

typedef struct Bytes20 {
  uint8_t bytes[20];
} Bytes20;

typedef struct RustSlice {
  const uint8_t *ptr;
  uintptr_t len;
} RustSlice;

typedef struct GoEvmApi {
  struct Bytes32 (*get_bytes32)(uintptr_t id, struct Bytes32 key, uint64_t *gas_cost);
  EvmApiStatus (*set_bytes32)(uintptr_t id,
                              struct Bytes32 key,
                              struct Bytes32 value,
                              uint64_t *gas_cost,
                              struct RustBytes *error);
  EvmApiStatus (*contract_call)(uintptr_t id,
                                struct Bytes20 contract,
                                struct RustSlice *calldata,
                                uint64_t *gas,
                                struct Bytes32 value,
                                uint32_t *return_data_len);
  EvmApiStatus (*delegate_call)(uintptr_t id,
                                struct Bytes20 contract,
                                struct RustSlice *calldata,
                                uint64_t *gas,
                                uint32_t *return_data_len);
  EvmApiStatus (*static_call)(uintptr_t id,
                              struct Bytes20 contract,
                              struct RustSlice *calldata,
                              uint64_t *gas,
                              uint32_t *return_data_len);
  EvmApiStatus (*create1)(uintptr_t id,
                          struct RustBytes *code,
                          struct Bytes32 endowment,
                          uint64_t *gas,
                          uint32_t *return_data_len);
  EvmApiStatus (*create2)(uintptr_t id,
                          struct RustBytes *code,
                          struct Bytes32 endowment,
                          struct Bytes32 salt,
                          uint64_t *gas,
                          uint32_t *return_data_len);
  void (*get_return_data)(uintptr_t id, struct RustBytes *output, uint32_t offset, uint32_t size);
  EvmApiStatus (*emit_log)(uintptr_t id, struct RustBytes *data, uint32_t topics);
  struct Bytes32 (*account_balance)(uintptr_t id, struct Bytes20 address, uint64_t *gas_cost);
  void (*account_code)(uintptr_t id,
                       struct RustBytes *code,
                       struct Bytes20 address,
                       uint32_t offset,
                       uint32_t size,
                       uint64_t *gas_cost);
  uint32_t (*account_code_size)(uintptr_t id, struct Bytes20 address, uint64_t *gas_cost);
  struct Bytes32 (*account_codehash)(uintptr_t id, struct Bytes20 address, uint64_t *gas_cost);
  uint64_t (*add_pages)(uintptr_t id, uint16_t pages);
  void (*capture_hostio)(uintptr_t id,
                         struct RustSlice *name,
                         struct RustSlice *args,
                         struct RustSlice *outs,
                         uint64_t start_ink,
                         uint64_t end_ink);
  uintptr_t id;
} GoEvmApi;

typedef struct EvmData {
  struct Bytes32 block_basefee;
  uint64_t chainid;
  struct Bytes20 block_coinbase;
  uint64_t block_gas_limit;
  uint64_t block_number;
  uint64_t block_timestamp;
  struct Bytes20 contract_address;
  struct Bytes20 msg_sender;
  struct Bytes32 msg_value;
  struct Bytes32 tx_gas_price;
  struct Bytes20 tx_origin;
  uint32_t reentrant;
  uint32_t return_data_len;
  bool tracing;
} EvmData;

typedef struct CByteArray {
  const uint8_t *ptr;
  uintptr_t len;
} CByteArray;

typedef struct GlobalState {
  struct Bytes32 bytes32_vals[GLOBAL_STATE_BYTES32_NUM];
  uint64_t u64_vals[GLOBAL_STATE_U64_NUM];
} GlobalState;

typedef struct ResolvedPreimage {
  uint8_t *ptr;
  intptr_t len;
} ResolvedPreimage;

typedef struct RustByteArray {
  uint8_t *ptr;
  uintptr_t len;
  uintptr_t capacity;
} RustByteArray;

/**
 * Instruments and "activates" a user wasm.
 *
 * The `output` is either the serialized asm & module pair or an error string.
 * Returns consensus info such as the module hash and footprint on success.
 *
 * Note that this operation costs gas and is limited by the amount supplied via the `gas` pointer.
 * The amount left is written back at the end of the call.
 *
 * # Safety
 *
 * `output`, `asm_len`, `module_hash`, `footprint`, and `gas` must not be null.
 */
UserOutcomeKind stylus_activate(struct GoSliceData wasm,
                                uint16_t page_limit,
                                uint16_t version,
                                bool debug,
                                struct RustBytes *output,
                                uintptr_t *asm_len,
                                struct Bytes32 *module_hash,
                                struct StylusData *stylus_data,
                                uint64_t *gas);

/**
 * Calls a compiled user program.
 *
 * # Safety
 *
 * `module` must represent a valid module produced from `stylus_activate`.
 * `output` and `gas` must not be null.
 */
UserOutcomeKind stylus_call(struct GoSliceData module,
                            struct GoSliceData calldata,
                            struct StylusConfig config,
                            struct GoEvmApi go_api,
                            struct EvmData evm_data,
                            uint32_t debug_chain,
                            struct RustBytes *output,
                            uint64_t *gas);

/**
 * Frees the vector. Does nothing when the vector is null.
 *
 * # Safety
 *
 * Must only be called once per vec.
 */
void stylus_drop_vec(struct RustBytes vec);

/**
 * Overwrites the bytes of the vector.
 *
 * # Safety
 *
 * `rust` must not be null.
 */
void stylus_vec_set_bytes(struct RustBytes *rust, struct GoSliceData data);

extern uint8_t wavm_caller_load8(uintptr_t ptr);

extern uint32_t wavm_caller_load32(uintptr_t ptr);

extern void wavm_caller_store8(uintptr_t ptr, uint8_t val);

extern void wavm_caller_store32(uintptr_t ptr, uint32_t val);

struct Machine *arbitrator_load_machine(const char *binary_path,
                                        const char *const *library_paths,
                                        intptr_t library_paths_size,
                                        uintptr_t debug_chain);

struct Machine *arbitrator_load_wavm_binary(const char *binary_path);

void arbitrator_free_machine(struct Machine *mach);

struct Machine *arbitrator_clone_machine(struct Machine *mach);

/**
 * Go doesn't have this functionality builtin for whatever reason. Uses relaxed ordering.
 */
void atomic_u8_store(uint8_t *ptr, uint8_t contents);

/**
 * Runs the machine while the condition variable is zero. May return early if num_steps is hit.
 * Returns a c string error (freeable with libc's free) on error, or nullptr on success.
 */
char *arbitrator_step(struct Machine *mach, uint64_t num_steps, const uint8_t *condition);

int arbitrator_add_inbox_message(struct Machine *mach,
                                 uint64_t inbox_identifier,
                                 uint64_t index,
                                 struct CByteArray data);

/**
 * Adds a user program to the machine's known set of wasms, compiling it into a link-able module.
 * Returns a c string error (freeable with libc's free) on compilation error, or nullptr on success.
 */
void arbitrator_add_user_wasm(struct Machine *mach,
                              const uint8_t *module,
                              uintptr_t module_len,
                              const struct Bytes32 *module_hash);

/**
 * Like arbitrator_step, but stops early if it hits a host io operation.
 * Returns a c string error (freeable with libc's free) on error, or nullptr on success.
 */
char *arbitrator_step_until_host_io(struct Machine *mach, const uint8_t *condition);

int arbitrator_serialize_state(const struct Machine *mach, const char *path);

int arbitrator_deserialize_and_replace_state(struct Machine *mach, const char *path);

uint64_t arbitrator_get_num_steps(const struct Machine *mach);

/**
 * Returns one of ARBITRATOR_MACHINE_STATUS_*
 */
uint8_t arbitrator_get_status(const struct Machine *mach);

struct GlobalState arbitrator_global_state(struct Machine *mach);

void arbitrator_set_global_state(struct Machine *mach, struct GlobalState gs);

void arbitrator_set_preimage_resolver(struct Machine *mach,
                                      struct ResolvedPreimage (*resolver)(uint64_t, const uint8_t*));

void arbitrator_set_context(struct Machine *mach, uint64_t context);

struct Bytes32 arbitrator_hash(struct Machine *mach);

struct Bytes32 arbitrator_module_root(struct Machine *mach);

struct RustByteArray arbitrator_gen_proof(struct Machine *mach);

void arbitrator_free_proof(struct RustByteArray proof);

#endif /* arbitrator_bindings */
